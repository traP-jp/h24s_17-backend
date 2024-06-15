package routes

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	e.GET("/hello/:name", hello)
	e.GET("/stand", GetTokenHandler)
}
