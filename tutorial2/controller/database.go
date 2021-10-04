package controller

import (
	"fmt"
	"go-echo-tutorial/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Server struct {
	DB     *gorm.DB
	Router *echo.Echo
}

func (server *Server) DBConnection(DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME string) {
	//host := "root:1111@tcp(127.0.0.1:3306)/rootdb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	host := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	server.DB, err = gorm.Open(mysql.Open(host), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	/*
		dbset, err := db.DB()

		if err != nil {
			log.Fatal(err)
		}

		dbset.SetMaxIdleConns(0)
		dbset.SetMaxOpenConns(5)
		dbset.SetConnMaxLifetime(time.Hour)
	*/

	server.DB.AutoMigrate(
		models.Shop{},
		models.User{},
	)

	setup, err := server.DB.DB()

	if err != nil {
		log.Fatal(err)
	}

	//server.DB.Migrator().CreateView()

	setup.SetMaxIdleConns(0)
	setup.SetMaxOpenConns(5)
	setup.SetConnMaxLifetime(time.Hour)

	server.Router = echo.New()
	server.ApiStore()
}

func (server *Server) AppLauncher(port string) {
	/*
		fmt.Printf("%s", port)

		server.Router = echo.New()

		server.Router.Logger.Fatal(server.Router.Start(port))
	*/

	log.Fatal(http.ListenAndServe(port, server.Router))
}
