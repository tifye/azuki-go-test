package api

import (
	a "azuki-server/azuki"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func registerRoutes(e *echo.Echo, _ *viper.Viper, logger *log.Logger) {
	e.GET("/", func(c echo.Context) error {
		logger.Debug("Azuki!")
		return c.String(http.StatusOK, "Azuki!")
	})

	start := time.Now()

	var counter atomic.Int64
	e.GET("/schema", handleSchema(logger.WithPrefix("schema")))
	e.GET("/nekopara/schema", handleNekoparaPage())
	e.POST("/trigger-a", handleTestTrigger(logger.WithPrefix("trigger-a")))
	e.POST("/trigger-b", handleTestTrigger(logger.WithPrefix("trigger-b")))
	e.GET("/chocola", handleReturnString("Chocola"))
	e.GET("/vanilla", handleReturnString("Vanilla"))
	e.POST("/counter/add", handleAdd(&counter))
	e.POST("/counter/subtract", handleSubtract(&counter))
	e.GET("/cinnamon", handleReturnStringFunc(func() string {
		since := time.Since(start)
		return fmt.Sprintf("%.0f", since.Seconds())
	}))
	e.GET("/coconut", handleReturnStringFunc(func() string {
		return fmt.Sprintf("%d", counter.Load())
	}))
}

func handleNekoparaPage() echo.HandlerFunc {
	// const base string = "http://192.168.18.175:8484"
	return func(c echo.Context) error {
		schema := a.NewSchema(
			a.Stack(a.Vertical).WithChildren(
				a.Image(a.Text("https://rare-gallery.com/mocahbig/77284-Chocola-Nekopara-Vanilla-Nekopara-Animal-Ears.png")),
				a.Label(a.Text("Chocola & Vanilla")),
			),
		)
		return c.JSON(http.StatusOK, schema)
	}
}

func handleSchema(logger *log.Logger) echo.HandlerFunc {
	const base string = "http://192.168.18.175:8484"
	const shigure string = "https://shigure-683956955842.europe-west1.run.app"
	const azukiImg string = "https://media.discordapp.net/attachments/1211775725628166186/1313884137291124870/c9942669afc4e7306009afaaf0e67f12-1.jpg?ex=67f5e435&is=67f492b5&hm=66ed2422872d8bdb787548e99aca0d343c305e6ca3738f8f28ef2282b6bdc3e5&=&format=webp&width=593&height=789"
	return func(c echo.Context) error {
		logger.Debug("Schema!")
		schema := a.NewSchema(
			a.Label(a.String("Welcome to Azuki!")),
			a.Stack(a.Horizontal).WithChildren(
				a.Stat(a.HTTPText(base+"/coconut")).
					WithTitle("Cinnamon").
					WithDescription("counter"),

				a.Stat(a.HTTPText(base+"/cinnamon").WithWatch(time.Second)).
					WithTitle("Shigure").
					WithPlace(a.PlaceCenter),

				a.Stat(a.HTTPText(base+"/cinnamon")).
					WithTitle("Coconut").
					WithDescription("counter").
					WithPlace(a.PlaceEnd),
			),
			a.Image(a.String(azukiImg)),
			a.Stack(a.Horizontal).WithChildren(
				a.Button(base+"/counter/add").
					WithText("+").
					WithInvalidatesTargets(base+"/coconut"),
				a.Button(base+"/counter/subtract").
					WithText("-").
					WithInvalidatesTargets(base+"/coconut"),
			),
			a.Button(base+"/trigger-b").
				WithText("Update").
				WithInvalidatesTargets(
					base+"/cinnamon",
					"https://www.youtube.com/watch?v=ixWYxt-C5qU&list=RDibbX6LMn_qs&index=3",
				),
			a.Stack(a.Horizontal).WithChildren(
				a.Button("").WithTextSource(a.HTTPText(base+"/chocola")),
				a.Button("").WithTextSource(a.HTTPText(base+"/vanilla")),
			),
			a.Stack(a.Vertical).WithChildren(
				a.Stack(a.Horizontal).WithGap(16).WithChildren(
					a.Label(a.HTTPText(shigure+"/activity").WitFieldpath("Author")).WithSize(16),
					a.Label(a.HTTPText(shigure+"/activity").WitFieldpath("Title")).WithSize(16),
				),
				a.Image(a.HTTPText(shigure+"/activity").WitFieldpath("ThumbnailUrl")),
			),
		)
		return c.JSON(http.StatusOK, schema)
	}
}

func handleTestTrigger(logger *log.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.Debug("Triggered!")
		return c.String(http.StatusOK, "Triggered!")
	}
}

func handleReturnString(str string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, str)
	}
}

func handleReturnStringFunc(f func() string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, f())
	}
}

func handleAdd(ai *atomic.Int64) echo.HandlerFunc {
	return func(c echo.Context) error {
		ai.Add(1)
		return c.NoContent(http.StatusOK)
	}
}

func handleSubtract(ai *atomic.Int64) echo.HandlerFunc {
	return func(c echo.Context) error {
		ai.Add(-1)
		return c.NoContent(http.StatusOK)
	}
}
