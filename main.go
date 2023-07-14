package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
// feature B

type User struct { 
	Id int `json:"id"` 
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
}

type BaseResponse struct {
	Message string `json:"message"` 
	Status bool  `json:"status"`
	Data interface{} `json:"data"` 
}

func main() {
	connectDatabase()
	e := echo.New()
	e.GET("/helloworld", HelloworldController)
	e.GET("/users", UserController)
	e.GET("/users/:id", DetailUserController)
	e.POST("/users", CreateUserController)
	e.Start(":8000")
}

func connectDatabase(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja6?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&User{})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	result := DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, BaseResponse{
			"Gagal memasukkan data", false, nil,
		})
	}
	return c.JSON(http.StatusCreated, BaseResponse{
		"Success memasukkan data", true, user,
	})
}

func DetailUserController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var user User = User{id, "Alterra", "alta@gmail.com", ""}
	return c.JSON(200, user)
}

func UserController(c echo.Context) error {
	var users []User 
	
	result := DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, BaseResponse{
			"Gagal mendapatkan data", false, nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		"Success mendapatkan data", true, users,
	})
}

func HelloworldController(c echo.Context) error {
	return c.JSON(200, "Hello World")
}


