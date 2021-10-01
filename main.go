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
	"gorm.io/gorm/schema"
)

type RequestBody struct {
	Message string `json:"message"`
}

type querybody struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Board struct {
	Idx          int       `gorm:"primary_key;auto_increment;not null" json:"idx"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Writer string `json:"writer"`
}

func DBConnection() {
	host := "root:1111@tcp(127.0.0.1:3306)/rootdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(host), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
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


func CreatePost(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		
		body, _ := ioutil.ReadAll(c.Request().Body)
		req := Board{}
		json.Unmarshal(body, &req)
		db.Model(&Board{}).Create(&req)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result" : "success",
		})
	}
}

func GetPost(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		idx, _ := strconv.Atoi(c.Param("idx"))
		board := Board{}
		db.First(&board, "idx = ?", idx)
		//fmt.Println(rows)
		return c.JSON(http.StatusOK, &board)
	}
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

	e2 := e.Group("/api/db")

	e2.POST("/c_post", CreatePost(db))
	e2.GET("/g_post/:idx", GetPost(db))


	e.Logger.Fatal(e.Start(":8081"))
}