package Entities

type Name1Entity struct {
	Id   int64
	Name string
}

func (ne Name1Entity) TableName() string {
	return "Names1"
}
