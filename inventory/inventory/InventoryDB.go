package inventory

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InventoryDB struct {
	db *gorm.DB
}

var inventoryrepository *InventoryDB

func InventoryDBInit() {
	var err error
	inventoryrepository = &InventoryDB{}
	inventoryrepository.db, err = gorm.Open(sqlite.Open("Inventory_table.db"), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	inventoryrepository.db.AutoMigrate(&Inventory{})

}

func InventoryRepository() *InventoryDB {
	return inventoryrepository
}

func (self *InventoryDB) save(entity interface{}) error {

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *InventoryDB) GetList() []Inventory {

	entities := []Inventory{}
	self.db.Find(&entities)

	return entities
}

func (self *InventoryDB) FindById(id int) (*Inventory, error) {
	entity := &Inventory{}
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

func (self *InventoryDB) Delete(entity *Inventory) error {
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *InventoryDB) Update(id int, params map[string]string) error {
	entity := &Inventory{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &Inventory{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}
