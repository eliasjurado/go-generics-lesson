package any

import "log"

func PrintList(list ...interface{}) {
	log.SetFlags(0)
	for _, v := range list {
		log.Printf("%+v\n", v)
	}
}

// any hace lo mismo que interface{}
func PrintListAny(list ...any) {
	log.SetFlags(0)
	for _, v := range list {
		log.Printf("%+v\n", v)
	}
}