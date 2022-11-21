package order

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderDB struct {
	db *gorm.DB
}

var orderrepository *OrderDB

func OrderDBInit() {
	var err error
	orderrepository = &OrderDB{}
	orderrepository.db, err = gorm.Open(sqlite.Open("Order_table.db"), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	orderrepository.db.AutoMigrate(&Order{})

}

func OrderRepository() *OrderDB {
	return orderrepository
}

func (self *OrderDB) save(entity interface{}) error {

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *OrderDB) GetList() []Order {

	entities := []Order{}
	self.db.Find(&entities)

	return entities
}

func (self *OrderDB) FindById(id int) (*Order, error) {
	entity := &Order{}
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

func (self *OrderDB) Delete(entity *Order) error {
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *OrderDB) Update(id int, params map[string]string) (*Order, error) { //TODO: template 수정
	entity := &Order{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &Order{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}
