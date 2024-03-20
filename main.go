package main

import (
	"generic/any"
	"generic/operator"
	"generic/restriction"
	"log"
)

func main() {
	//any
	any.PrintList("Elias", "Jurado", "Santos")
	any.PrintListAny(0, 1, 14, 35.5)

	//restricciones
	log.Printf("%+v\n", restriction.Sum(4, 5, 6))
	log.Printf("%+v\n", restriction.Sum(4.2, 5.34, 6.0035))

	log.Printf("%+v\n", restriction.SumAproximacion(4, 5, 6))
	log.Printf("%+v\n", restriction.SumAproximacion(4.2, 5.34, 6.0035))

	log.Printf("%+v\n", restriction.SumConstraints(4, 5, 6))
	log.Printf("%+v\n", restriction.SumConstraints(4.2, 5.34, 6.0035))

	//operadores
	strings := []string{"a", "b", "c", "d", "e"}
	numbers := []int{1, 2, 3, 4, 5, 6}

	log.Printf("%+v\n", operator.Includes(strings, "a"))
	log.Printf("%+v\n", operator.Includes(strings, "c"))

	log.Printf("%+v\n", operator.Includes(numbers, 1))
	log.Printf("%+v\n", operator.Includes(numbers, 5))

	log.Printf("%+v\n", operator.Filter(numbers, func(value int) bool { return value > 3 }))
	log.Printf("%+v\n", operator.Filter(strings, func(value string) bool { return value > "b" }))

}
