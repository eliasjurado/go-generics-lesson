package product

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}