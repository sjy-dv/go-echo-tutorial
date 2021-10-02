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
	Age  int    `json:"age"`
}

type Board struct {
	gorm.Model
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Writer string `json:"writer"`
}

type BoardBody struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
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

	db.AutoMigrate(&Board{})

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
		"result": "hello world~!!",
	})
}

func BodyPostTest(c echo.Context) error {

	body, _ := ioutil.ReadAll(c.Request().Body)
	req := RequestBody{}

	json.Unmarshal(body, &req)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": req.Message,
	})
}

func EachQueryTest(c echo.Context) error {

	first_query := c.QueryParam("name")
	second_query, _ := strconv.Atoi(c.QueryParam("age"))

	return c.JSON(http.StatusOK, querybody{
		Name: first_query,
		Age:  second_query,
	})
}

func MulQueryTest(c echo.Context) error {
	req := c.QueryParams()
	first_query := req["name"][0]
	second_query, _ := strconv.Atoi(req["age"][0])

	return c.JSON(http.StatusOK, querybody{
		Name: first_query,
		Age:  second_query,
	})
}

func ParamsTest(c echo.Context) error {
	req := c.Param("num")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": fmt.Sprintf("받은 값 : %s", req),
	})
}

func CreatePost(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		body, _ := ioutil.ReadAll(c.Request().Body)
		req := Board{}
		json.Unmarshal(body, &req)
		err := db.Debug().Model(&Board{}).Create(&req).Error
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": "fail",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": "success",
		})
	}
}

func GetPostParam(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		board := Board{}
		err := db.Debug().First(&board, "id = ?", id).Error
		//fmt.Println(rows)
		if err != nil {
			return c.JSON(http.StatusOK, &Board{})
		}
		return c.JSON(http.StatusOK, &board)
	}
}

func GetPostQuery(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		board := Board{}
		err := db.Debug().First(&board, "id = ?", id).Error
		if err != nil {
			return c.JSON(http.StatusOK, &Board{})
		}
		return c.JSON(http.StatusOK, &board)
	}
}

func GetAllPost(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		offset := 0
		if page > 1 {
			offset = 10 * (page - 1)
		}
		boards := []Board{}
		err := db.Debug().Model(&Board{}).Limit(10).Offset(offset).Scan(&boards).Error

		if err != nil {
			return c.JSON(http.StatusOK, &[]Board{})
		}
		return c.JSON(http.StatusOK, &boards)
	}
}

func UpdatePost(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		body, _ := ioutil.ReadAll(c.Request().Body)
		req := BoardBody{}
		json.Unmarshal(body, &req)

		err := db.Debug().Model(&Board{}).Where("id = ?", req.Id).Updates(Board{
			Title: req.Title,
			Desc:  req.Desc,
		}).Error
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": "fail",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": "success",
		})
	}
}

func DeletePost(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		body, _ := ioutil.ReadAll(c.Request().Body)
		req := BoardBody{}
		json.Unmarshal(body, &req)

		err := db.Debug().Delete(&Board{}, req.Id)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": "fail",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": "success",
		})
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
	e2.GET("/g_post/:id", GetPostParam(db))
	e2.GET("/g_qpost", GetPostQuery(db))
	e2.GET("/g_all", GetAllPost(db))
	e2.POST("/u_post", UpdatePost(db))
	e2.POST("/d_post", DeletePost(db))
	e.Logger.Fatal(e.Start(":8081"))
}
