package view

type UserShopInfoView struct {
	UserUserid   string
	UserUsername string
	ShopShopName string
	ShopShopInfo string
}

func (v UserShopInfoView) ViewName() string {
	return "view_user_shop_info"
}
