package models

import (
	"go-echo-tutorial/models/view"
	"go-echo-tutorial/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) FindUser(db *gorm.DB) (*[]User, error) {

	users := []User{}

	err := db.Debug().Model(&User{}).Scan(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func (u *User) CreateUser(db *gorm.DB, userid, username, password string) error {

	hashpassword, err := utils.BcryptHashSync(password)
	if err != nil {
		return err
	}
	u.Userid = userid
	u.Username = username
	u.Password = string(hashpassword)

	err = db.Debug().Create(&u).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *User) LoginUser(db *gorm.DB, userid, password string) (string, error) {

	err := db.Debug().Model(&User{}).Take(&u).Error

	if err != nil {
		return "", err
	}
	err = utils.BcryptCompareSync(u.Password, password)

	if err != nil {
		return "", err
	}

	token, err2 := utils.CreateToken(u.Userid)

	if err2 != nil {
		return "", err2
	}

	return token, nil
}

func (s *Shop) CheckShop(db *gorm.DB, userid string) bool {

	err := db.Debug().Model(&Shop{}).Take(&s).Error

	if err != nil {
		return false
	}

	return true
}

func (u *User) UserInfo(db *gorm.DB, userid string) (*User, error) {

	err := db.Debug().Model(&User{}).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func UserShopInfo(db *gorm.DB, userid string) (*view.UserShopInfoView, error) {

	queryResult := view.UserShopInfoView{}
	err := db.Debug().Table(view.UserShopInfoView{}.ViewName()).Where("user_userid = ?", userid).Find(&queryResult).Error
	if err != nil {
		return &view.UserShopInfoView{}, err
	}

	return &queryResult, nil
}
