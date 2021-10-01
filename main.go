package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RequestBody struct {
	Message string `json:"message"`
}

type querybody struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func DBConnection() {
	host := "root:1111@tcp(127.0.0.1:3306)/rootdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(host), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	handlerouting(db)

	dbset, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	dbset.SetMaxIdleConns(0)
	dbset.SetMaxOpenConns(5)
	dbset.SetConnMaxLifetime(time.Hour)
}

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result" : "hello world~!!",
	})
}

func BodyPostTest(c echo.Context) error {

	body, _ := ioutil.ReadAll(c.Request().Body)
	req := RequestBody{}

	json.Unmarshal(body, &req)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result" : req.Message,
	})
}

func EachQueryTest(c echo.Context) error {

	first_query := c.QueryParam("name")
	second_query, _ := strconv.Atoi(c.QueryParam("age"))

	return c.JSON(http.StatusOK, querybody{
		Name: first_query,
		Age: second_query,
	})
}

func MulQueryTest(c echo.Context) error {
	req := c.QueryParams()
	first_query := req["name"][0]
	second_query, _ := strconv.Atoi(req["age"][0]) 

	return c.JSON(http.StatusOK, querybody{
		Name: first_query,
		Age: second_query,
	})
}

func ParamsTest(c echo.Context) error {
	req := c.Param("num")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result" : fmt.Sprintf("받은 값 : %s", req),
	})
}

func main() {
	DBConnection()	
}

func handlerouting(db *gorm.DB) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/", HelloWorld)
	e.POST("/bodypost", BodyPostTest)
	e.GET("/q1", EachQueryTest)
	e.GET("/q2", MulQueryTest)
	e.GET("/p/:num", ParamsTest)
	e.Logger.Fatal(e.Start(":8081"))
}