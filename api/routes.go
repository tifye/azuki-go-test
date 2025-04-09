package api

import (
	a "azuki-server/azuki"
	"fmt"
	"net/http"
	"strings"
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
	links := map[string]string{
		"chocola":  "https://rare-gallery.com/mocahbig/77284-Chocola-Nekopara-Vanilla-Nekopara-Animal-Ears.png",
		"vanilla":  "https://rare-gallery.com/mocahbig/77284-Chocola-Nekopara-Vanilla-Nekopara-Animal-Ears.png",
		"maple":    "https://rare-gallery.com/thumbnail/71362-Maple-Nekopara-Cinnamon-NekoparaNEKOPARA-Vol..jpg",
		"cinnamon": "https://rare-gallery.com/thumbnail/71362-Maple-Nekopara-Cinnamon-NekoparaNEKOPARA-Vol..jpg",
		"coconut":  "https://w0.peakpx.com/wallpaper/253/123/HD-wallpaper-nekopara-nekopara-vol-2-azuki-nekopara-coconut-nekopara-heterochromia.jpg",
		"azuki":    "https://w0.peakpx.com/wallpaper/253/123/HD-wallpaper-nekopara-nekopara-vol-2-azuki-nekopara-coconut-nekopara-heterochromia.jpg",
		"shigure":  "https://preview.redd.it/daily-neko-girl-day-50-3-shigure-looked-stunning-in-this-v0-0abxsb1rwo0e1.jpeg?auto=webp&s=35f2a6116e4baa667f62f6c2e5d7c2b02de0c6cd",
	}
	type request struct {
		Input string `query:"input"`
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		input := strings.ToLower(req.Input)
		link, ok := links[input]
		if !ok {
			schema := a.NewSchema(
				a.Label(a.String("Not found ):")),
			)
			return c.JSON(http.StatusOK, schema)
		}

		schema := a.NewSchema(
			a.Stack(a.Vertical).WithChildren(
				a.Image(a.Text(link)),
				a.Label(a.Text(strings.ToTitle(input))),
			),
		)
		return c.JSON(http.StatusOK, schema)
	}
}

func handleSchema(logger *log.Logger) echo.HandlerFunc {
	const base string = "http://192.168.18.175:8484"
	const shigure string = "https://shigure-683956955842.europe-west1.run.app"
	return func(c echo.Context) error {
		logger.Debug("Schema!")
		schema := a.NewSchema(
			a.Label(a.String("Welcome to Azuki!")),
			a.TextInput(),
			a.Stack(a.Vertical).WithChildrenKey("search"),
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
