package restriction

func Sum[T int | float64](values ...T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// usando la virgulilla ~ puedo usar todos los datos que derivan del tipo de dato indicado
func SumAproximacion[T ~int | ~float64](values ...T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}
