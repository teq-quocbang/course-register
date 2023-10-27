package http

import (
	"net/http"
	"regexp"

	echoSentry "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/delivery/http/account"
	"github.com/teq-quocbang/course-register/delivery/http/auth"
	"github.com/teq-quocbang/course-register/delivery/http/class"
	"github.com/teq-quocbang/course-register/delivery/http/course"
	"github.com/teq-quocbang/course-register/delivery/http/example"
	"github.com/teq-quocbang/course-register/delivery/http/healthcheck"
	"github.com/teq-quocbang/course-register/delivery/http/semester"
	"github.com/teq-quocbang/course-register/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e         = echo.New()
		loggerCfg = middleware.DefaultLoggerConfig
	)

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health-check"
	}

	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echoSentry.New(echoSentry.Options{Repanic: true}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOriginFunc: func(origin string) (bool, error) {
			return regexp.MatchString(
				`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
				origin,
			)
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// Health check
	healthcheck.Init(e.Group("/health-check"))

	// API docs
	if !config.GetConfig().Stage.IsProd() {
		e.GET("/docs/*", echoSwagger.WrapHandler)
	}

	// APIs
	api := e.Group("/api")
	example.Init(api.Group("/examples"), useCase)
	account.Init(api.Group("/user"), useCase)
	semester.Init(api.Group("/semester", auth.Auth), useCase)
	class.Init(api.Group("/class", auth.Auth), useCase)
	course.Init(api.Group("/course", auth.Auth), useCase)

	return e
}
