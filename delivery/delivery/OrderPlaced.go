package delivery

import (
	"time"
)

type OrderPlaced struct {
	EventType  string `json:"eventType" type:"string"`
	TimeStamp  string `json:"timeStamp" type:"string"`
	Id         int    `json:"id" type:"int"`
	ProductId  string `json:"productId" type:"string"`
	Qty        int    `json:"qty" type:"int"`
	CustomerId string `json:"customerId" type:"string"`
	Amount     int    `json:"amount" type:""` // External Event 에 대하여 type 중 번역이 누락된듯? 이것만 왜 안나온??
	Status     string `json:"status" type:"string"`
	Address    string `json:"address" type:"string"`
}

func NewOrderPlaced() *OrderPlaced {
	event := &OrderPlaced{EventType: "OrderPlaced", TimeStamp: time.Now().String()}

	return event
}
