package inventory

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverOrderPlaced_DecreaseStock(data map[string]interface{}){
	
	event := NewOrderPlaced()
	mapstructure.Decode(data,&event)
	inventory := &Inventory{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	inventoryrepository.save(inventory)

	// Sample Logic //
	inventory.DecreaseStock(event);
}

func wheneverOrderCancelled_IncreaseStock(data map[string]interface{}){
	
	event := NewOrderCancelled()
	mapstructure.Decode(data,&event)
	inventory := &Inventory{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	inventoryrepository.save(inventory)

	// Sample Logic //
	inventory.IncreaseStock(event);
}

