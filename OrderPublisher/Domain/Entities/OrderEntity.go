package Entities

type OrderEntity struct {
	Id    uint64
	Price uint
	Title string
}

func (ne OrderEntity) TableName() string {
	return "orders"
}
