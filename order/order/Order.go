package order

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"order/external"
)

type Order struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	ProductId string `json:"productId"`
	Qty int `json:"qty"`
	CustomerId string `json:"customerId"`
	Amount int
	Status string `json:"status"`
	Address string `json:"address"`

}

func (self *Order) onPostPersist() (err error){
	orderPlaced := NewOrderPlaced()
	model.Copy(orderPlaced, self)

	Publish(orderPlaced)

	return nil
}
func (self *Order) onPrePersist() (err error){

	// Get request from Inventory
	inventory := &external.Inventory{}
	resp := external.GetInventory(inventory.Id)
	fmt.Println(resp)

	return nil
}
func (self *Order) onPreRemove() (err error){
	orderCancelled := NewOrderCancelled()
	model.Copy(orderCancelled, self)

	Publish(orderCancelled)

	return nil
}
func (self *Order) onPreUpdate() (err error){ return nil }
func (self *Order) onPostUpdate() (err error){ return nil }
func (self *Order) onPostRemove() (err error){ return nil }


