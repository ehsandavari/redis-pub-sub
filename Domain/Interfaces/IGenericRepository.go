package Interfaces

type IGenericRepository[T any] interface {
	Find() (T, error)
	Add(model *T) (T, error)
}
