package delivery

import (
	"time"
)

type DeliveryStarted struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	Address string `json:"address" type:"string"` 
	CustomerId string `json:"customerId" type:"string"` 
	Quantity int `json:"quantity" type:"int"` 
	OrderId int `json:"orderId" type:"int"` 
	
}

func NewDeliveryStarted() *DeliveryStarted{
	event := &DeliveryStarted{EventType:"DeliveryStarted", TimeStamp:time.Now().String()}

	return event
}
