create view view_user_shop_info
as
    select user.userid as user_userid,
        user.username as user_username,
        shop.shop_name as shop_shop_name,
        shop.shop_info as shop_shop_info
    from user inner join shop
        on user.userid = shop.userid;