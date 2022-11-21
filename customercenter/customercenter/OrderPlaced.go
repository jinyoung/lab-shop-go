package customercenter

import (
	"time"
)

type OrderPlaced struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"`
	ProductId string `json:"productId" type:"string"`
	Qty int `json:"qty" type:"int"`
	CustomerId string `json:"customerId" type:"string"`
	Amount  `json:"amount" type:""`
	Status string `json:"status" type:"string"`
	Address string `json:"address" type:"string"`
}

func NewOrderPlaced() *OrderPlaced{
	event := &OrderPlaced{EventType:"OrderPlaced", TimeStamp:time.Now().String()}
	return event
}
