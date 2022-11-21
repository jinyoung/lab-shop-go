package delivery

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverOrderPlaced_AddToDeliveryList(data map[string]interface{}){
	
	event := NewOrderPlaced()
	mapstructure.Decode(data,&event)
	delivery := &Delivery{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	deliveryrepository.save(delivery)

	// Sample Logic //
	delivery.AddToDeliveryList(event);
}

