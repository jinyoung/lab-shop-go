package customercenter

import (
	"github.com/mitchellh/mapstructure"
	"log"
)

func whenOrderPlaced_then_CREATE_1 (inputEvent map[string]interface{}) {

	orderPlaced := NewOrderPlaced()
	mapstructure.Decode(inputEvent, &orderPlaced)

	myPage := &MyPage{}
	myPage.OrderId = orderPlaced.Id
	myPage.ProductId = orderPlaced.ProductId
	myPage.OrderStatus = orderPlaced.직접입력

	// view 레파지 토리에 save
	repository := MyPageRepository()
	err := repository.save(myPage)
	if err != nil {
		// TODO error control
		log.Printf("Create error: %v \n", err)
	}
}

func whenDeliveryStarted_then_UPDATE_1 (inputEvent map[string]interface{}) {

	deliveryStarted := NewDeliveryStarted()
	mapstructure.Decode(inputEvent,&deliveryStarted)

	var myPages []MyPage
	repository := MyPageRepository()
	// FIXME geom lib define snake_case as column name (eg: user_id), so if your query column is 'userId' then change 'user_id'
	err := repository.db.Where("orderId = ?", deliveryStarted.OrderId).Find(&myPages).Error
	if err != nil {
		// TODO error control
		log.Printf("Select error: %v \n", err)
	}
	for _, viewEntity := range myPages {
		viewEntity.DeliveryStatus = deliveryStarted.직접입력
		err1 := repository.db.Updates(viewEntity).Error
		if err1 != nil {
			// TODO error control
			log.Printf("Update error: %v \n", err1)
		}
	}
}

