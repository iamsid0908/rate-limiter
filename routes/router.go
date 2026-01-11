package route

import (
	"my-echo-app/middleware"
	"net/http"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
)

func InitHttp() *echo.Echo {
	app := App()
	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"https://book-finder0908sid.netlify.app", "http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// Apply your custom rate limiter middleware
	e.Use(middleware.RateLimitMiddleware(app.RedisClient))

	v1Routes(e.Group("/v1"), app)
	return e
}
