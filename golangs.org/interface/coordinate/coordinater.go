package main

import "fmt"

type coordinate struct {
	D, M, S float64
	H       rune
}

// String форматирует координаты DMS
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.D, c.M, c.S, c.H)
}

type location struct {
	Lat, Long coordinate
}

// String форматирует location  с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Long)
}

func main() {
	elysium := location{
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium)
}
