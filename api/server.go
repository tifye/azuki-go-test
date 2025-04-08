package api

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"github.com/spf13/viper"
)

func NewServer(config *viper.Viper, logger *log.Logger) *http.Server {
	e := echo.New()

	server := http.Server{
		Handler:           e,
		ReadTimeout:       15 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		ErrorLog:          logger.StandardLog(),
	}

	rlimiterConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      rate.Limit(5),
				Burst:     15,
				ExpiresIn: 3 * time.Minute,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			return ctx.RealIP(), nil
		},
		DenyHandler: func(ctx echo.Context, identifier string, err error) error {
			logger.Warn("ratelimiting", "identifier", identifier, "err", err)
			return ctx.NoContent(http.StatusTooManyRequests)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(rlimiterConfig))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestedWith, echo.HeaderAuthorization},
	}))

	eLoggerConfig := middleware.DefaultLoggerConfig
	eLoggerConfig.Output = logger.StandardLog().Writer()
	e.Use(
		middleware.LoggerWithConfig(eLoggerConfig),
	)

	registerRoutes(e, config, logger)

	return &server
}
