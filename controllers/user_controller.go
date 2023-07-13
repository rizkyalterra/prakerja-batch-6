package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: user,
	})
}

func DetailUserController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var user models.User = models.User{Id: id, Name: "Alterra", Email: "alta@gmail.com", Address: ""}
	return c.JSON(200, user)
}

func UserController(c echo.Context) error {
	var users []models.User 
	
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: users,
	})
}