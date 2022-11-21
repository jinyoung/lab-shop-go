package delivery
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type DeliveryDB struct{
	db *gorm.DB
}

var deliveryrepository *DeliveryDB

func DeliveryDBInit() {
	var err error
	deliveryrepository = &DeliveryDB{}
	deliveryrepository.db, err = gorm.Open(sqlite.Open("Delivery_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	deliveryrepository.db.AutoMigrate(&Delivery{})

}

func DeliveryRepository() *DeliveryDB {
	return deliveryrepository
}

func (self *DeliveryDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *DeliveryDB)GetList() []Delivery{
	
	entities := []Delivery{}
	self.db.Find(&entities)

	return entities
}

func (self *DeliveryDB)FindById(id int) (*Delivery, error){
	entity := &Delivery{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		return entity, nil
	}
}

func (self *DeliveryDB) Delete(entity *Delivery) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *DeliveryDB) Update(id int, params map[string]string) error{
	entity := &Delivery{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &Delivery{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}