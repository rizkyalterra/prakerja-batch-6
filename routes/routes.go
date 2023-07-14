package routes

import (
	"os"
	"prakerja6/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) *echo.Echo{
	e.Use(middleware.Logger())
	e.POST("/login", controllers.LoginController)
	e.Use(echojwt.JWT([]byte(os.Getenv("KEY_ENCRYPTOPN"))))
	e.GET("/users", controllers.UserController)
	e.GET("/users/:id", controllers.DetailUserController)
	e.POST("/users", controllers.CreateUserController)
	
	return e
}