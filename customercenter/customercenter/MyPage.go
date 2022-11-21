package customercenter

type MyPage struct {
	OrderId int `gorm:"primaryKey" json:"id" type:"int"`
	ProductId string `json:"productId" type:"string"`
	DeliveryStatus string `json:"deliveryStatus" type:"string"`
	OrderStatus string `json:"orderStatus" type:"string"`
}