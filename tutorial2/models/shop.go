package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Userid   string `json:"userid"`
	Shopname string `json:"shopname"`
	ShopInfo string `json:"shop_info"`
}

func (s *Shop) CreateShop(db *gorm.DB, userid string) error {

	s.Userid = userid

	err := db.Debug().Create(&s).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *Shop) UpdateShopInfo(db *gorm.DB, userid string) error {

	err := db.Debug().Where("userid = ?", userid).Updates(Shop{
		Shopname: s.Shopname,
		ShopInfo: s.ShopInfo,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
