package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (s *Server) ApiStore() {

	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	s.Router.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	userApi := s.Router.Group("/api/user")
	userApi.GET("/alluser", s.AllUser)
	userApi.POST("/signup", s.SignUp)
	userApi.POST("/signin", s.SingIn)
	userApi.POST("/updateso", s.UpdateShopOwner)
	userApi.GET("/myinfo", s.MyInform)

	shopApi := s.Router.Group("/api/shop")
	shopApi.POST("/updateinfo", s.UpdateShopInfo)
}
