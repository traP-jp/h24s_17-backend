package routes

import "github.com/labstack/echo/v4"

func (s *State) SetupRoutes(e *echo.Echo) {
	e.GET("/hello/:name", s.HelloHandler)
	e.POST("/checkin", s.CheckinHandler)
}
