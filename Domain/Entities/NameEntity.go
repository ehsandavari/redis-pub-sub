package Entities

type NameEntity struct {
	Id   int64
	Name string
}

func (ne NameEntity) TableName() string {
	return "Names"
}
