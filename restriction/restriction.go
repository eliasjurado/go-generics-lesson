package restriction

import "golang.org/x/exp/constraints"

func Sum[T int | float64](values ...T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// usando la virgulilla ~ puedo usar todos los datos que derivan del tipo de dato indicado
func SumAproximacion[T Numbers](values ...T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// creamos nuestra propia restriccion que podemos ir progresivamente ampliando sin necesidad de modificar nuestra función
type Numbers interface {
	~int | ~float32 | ~float64 | uint
}

// podemos reutilizar constraints de otro paquete
func SumConstraints[T constraints.Integer | constraints.Float](values ...T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}
