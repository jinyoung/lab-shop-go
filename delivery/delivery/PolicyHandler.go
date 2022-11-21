package delivery

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverOrderPlaced_AddToDeliveryList(data map[string]interface{}) { //TODO: template has been changed

	event := NewOrderPlaced()
	mapstructure.Decode(data, &event)

	AddToDeliveryList(event)
}
