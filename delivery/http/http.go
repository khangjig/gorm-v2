package http

import (
	"fmt"
	"gorm-v2/delivery/http/socket"
	"gorm-v2/delivery/http/test"
	"gorm-v2/usecase"
	"log"
	"net/http"
	"os"

	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	loggerCfg := middleware.DefaultLoggerConfig

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.RequestURI() == "/health_check"
	}

	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())

	if os.Getenv("IS_LOCAL") != "true" {
		e.Use(sentryEcho.New(sentryEcho.Options{
			Repanic: true,
		}))
	}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if c.Request().URL.RequestURI() != "/health_check" {
			request := fmt.Sprintf("%s", reqBody)

			if len(request) > 0 {
				log.Printf("%s", request)
			}
			log.Printf("%s", resBody)
		}
	}))

	e.GET("/health_check", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "ok"})
	})

	api := e.Group("/api")
	socket.Init(api, useCase)
	test.Init(api, useCase)

	return e
}
