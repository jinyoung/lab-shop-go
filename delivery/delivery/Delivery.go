package delivery

import (
	"gopkg.in/jeevatkm/go-model.v1"

	"gorm.io/gorm"
	//"delivery/external"
)

type Delivery struct {
	gorm.Model
	Id         int    `gorm:"primaryKey" json:"id" type:"int"`
	Address    string `json:"address"`
	CustomerId string `json:"customerId"`
	Quantity   int    `json:"quantity"`
	OrderId    int    `json:"orderId"`
}

func (self *Delivery) onPostPersist() (err error) {
	deliveryStarted := NewDeliveryStarted()
	model.Copy(deliveryStarted, self)

	Publish(deliveryStarted)

	return nil
}
func (self *Delivery) onPrePersist() (err error) { return nil }
func (self *Delivery) onPreUpdate() (err error)  { return nil }
func (self *Delivery) onPostUpdate() (err error) { return nil }
func (self *Delivery) onPreRemove() (err error)  { return nil }
func (self *Delivery) onPostRemove() (err error) { return nil }

func AddToDeliveryList(orderPlaced *OrderPlaced) {
	/** Example 1:  new item
	delivery := &Delivery{}
	deliveryrepository.save(delivery)

	*/

	// event := NewOrderPlaced()
	// mapstructure.Decode(data,&event)
	delivery := &Delivery{}
	delivery.OrderId = orderPlaced.Id

	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	deliveryrepository.save(delivery)

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(orderPlaced.id, 10, 64)
	point, err := PointRepository().FindById(int(id))
	if err != nil {

	}
	*/
}
