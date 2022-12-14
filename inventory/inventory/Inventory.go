package inventory

import (
	"gorm.io/gorm"
	//"inventory/external"
)

type Inventory struct {
	gorm.Model
	Id    int `gorm:"primaryKey" json:"id" type:"int"`
	Stock int `json:"stock"`
}

func (self *Inventory) onPostPersist() (err error) {

	return nil
}
func (self *Inventory) onPrePersist() (err error) { return nil }
func (self *Inventory) onPreUpdate() (err error)  { return nil }
func (self *Inventory) onPostUpdate() (err error) { return nil }
func (self *Inventory) onPreRemove() (err error)  { return nil }
func (self *Inventory) onPostRemove() (err error) { return nil }

func DecreaseStock(orderPlaced *OrderPlaced) {
	/** Example 1:  new item
	inventory := &Inventory{}
	inventoryrepository.save(inventory)

	*/

	/** Example 2:  finding and process	*/

	inventory, err := inventoryrepository.FindById(orderPlaced.Id)
	if err != nil {
		inventory.Stock = inventory.Stock - 1
		inventoryrepository.save(inventory)
	}

}
func IncreaseStock(orderCancelled *OrderCancelled) {
	/** Example 1:  new item
	inventory := &Inventory{}
	inventoryrepository.save(inventory)

	*/

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(orderCancelled.id, 10, 64)
	point, err := PointRepository().FindById(int(id))
	if err != nil {

	}
	*/
}
