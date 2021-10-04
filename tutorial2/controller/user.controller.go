package controller

import (
	"encoding/json"
	"go-echo-tutorial/models"
	"go-echo-tutorial/utils"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) AllUser(c echo.Context) error {

	u := models.User{}

	rows, err := u.FindUser(s.DB)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, rows)
}

func (s *Server) SignUp(c echo.Context) error {

	body, _ := ioutil.ReadAll(c.Request().Body)

	u := models.User{}
	_ = json.Unmarshal(body, &u)

	err := u.CreateUser(s.DB, u.Userid, u.Username, u.Password)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}

func (s *Server) SingIn(c echo.Context) error {

	body, _ := ioutil.ReadAll(c.Request().Body)

	u := models.User{}

	_ = json.Unmarshal(body, &u)
	token, err := u.LoginUser(s.DB, u.Userid, u.Password)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": "data_error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (s *Server) UpdateShopOwner(c echo.Context) error {

	userid := utils.VerifyTokenHeader(c)
	sh := models.Shop{}

	err := sh.CreateShop(s.DB, userid)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}

func (s *Server) UpdateShopInfo(c echo.Context) error {

	userid := utils.VerifyTokenHeader(c)
	body, _ := ioutil.ReadAll(c.Request().Body)
	sh := models.Shop{}
	_ = json.Unmarshal(body, &sh)
	err := sh.UpdateShopInfo(s.DB, userid)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"result": "fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}

func (s *Server) MyInform(c echo.Context) error {

	userid := utils.VerifyTokenHeader(c)
	sh := models.Shop{}
	u := models.User{}
	check := sh.CheckShop(s.DB, userid)

	if check == false {
		rows, err := u.UserInfo(s.DB, userid)

		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": "fail",
			})
		}
		return c.JSON(http.StatusOK, rows)
	} else if check == true {
		rows, err := models.UserShopInfo(s.DB, userid)

		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": "fail",
			})
		}
		return c.JSON(http.StatusOK, rows)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"result": "fatal error",
		})
	}
}
