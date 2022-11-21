package customercenter
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type MyPageDB struct{
	db *gorm.DB
}

var myPagerepository *MyPageDB

func MyPageDBInit() {
	var err error
	myPagerepository = &MyPageDB{}
	myPagerepository.db, err = gorm.Open(sqlite.Open("MyPage_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	myPagerepository.db.AutoMigrate(&MyPage{})

}

func MyPageRepository() *MyPageDB {
	return myPagerepository
}


func (self *MyPageDB)save(entity interface{}) error {

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *MyPageDB)GetList() []MyPage{

	entities := []MyPage{}
	self.db.Find(&entities)

	return entities
}

func (self *MyPageDB)FindById(id int) (*MyPage, error){
	entity := &MyPage{}
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

func (self *MyPageDB) Delete(entity *MyPage) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *MyPageDB) Update(id int, params map[string]string) (*MyPage, error){
	entity := &MyPage{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &MyPage{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}