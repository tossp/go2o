/**
 * Copyright 2013 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2013-12-22 21:56
 * description :
 * history :
 */

package dto

type ShoppingCart struct {
	Id         int         `json:"-"`
	CartKey    string      `json:"key"`
	BuyerId    int         `json:"buyer"`
	Summary    string      `json:"summary"`
	UpdateTime int64       `json:"update_time"`
	Items      []*CartItem `json:"items"`
	TotalNum   int         `json:"total_num"` // 总数量
	TotalFee   float32     `json:"total"`
	OrderFee   float32     `json:"fee"`
}

type CartItem struct {
	GoodsId    int     `json:"id"`
	GoodsName  string  `json:"name"`
	GoodsNo    string  `json:"no"`
	SmallTitle string  `json:"title"`
	GoodsImage string  `json:"image"`
	Num        int     `json:"num"`
	Price      float32 `json:"price"`
	SalePrice  float32 `json:"sale_price"`
}
