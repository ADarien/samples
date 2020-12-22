package main

import "fmt"

// location широта и долгота
type location struct {
	Lat, Long float64
}

// String форматирует location с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Long)
}

func main() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}
