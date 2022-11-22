package inventory

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverOrderPlaced_DecreaseStock(data map[string]interface{}) {

	event := NewOrderPlaced()
	mapstructure.Decode(data, &event)

	// Sample Logic //
	DecreaseStock(event)
}

func wheneverOrderCancelled_IncreaseStock(data map[string]interface{}) {

	event := NewOrderCancelled()
	mapstructure.Decode(data, &event)
	IncreaseStock(event)
}
