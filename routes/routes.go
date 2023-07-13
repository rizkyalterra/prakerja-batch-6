package routes

import (
	"prakerja6/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo{
	e.GET("/users", controllers.UserController)
	e.GET("/users/:id", controllers.DetailUserController)
	e.POST("/users", controllers.CreateUserController)
	return e
}