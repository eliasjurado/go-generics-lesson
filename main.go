package main

import (
	"generic/any"
	"generic/restriction"
	"log"
)

func main() {
	//any
	any.PrintList("Elias", "Jurado", "Santos")
	any.PrintListAny(0, 1, 14, 35.5)

	//restricciones
	log.Printf("%+v\n", restriction.Sum(4, 5, 6))
	log.Printf("%+v\n", restriction.Sum(4.2,5.34,6.0035))

	log.Printf("%+v\n", restriction.SumAproximacion(4, 5, 6))
	log.Printf("%+v\n", restriction.SumAproximacion(4.2,5.34,6.0035))
}
