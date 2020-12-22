package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// coordinate в градусах, минутах и секундах в сфере N/S/E/W
type coordinate struct {
	D, M, S float64
	H       rune
}

// String форматирует координаты DMS
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.D, c.M, c.S, c.H)
}

// Decimal конвертирует координаты d/m/s в десятичные градусы
func (c coordinate) Decimal() float64 {
	sign := 1.0
	switch c.H {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		//объявление
		DD  float64 `json:"deciaml"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		// заполнение
		DD:  c.Decimal(),
		DMS: c.String(),
		D:   c.D,
		M:   c.M,
		S:   c.S,
		H:   string(c.H),
	})
}

// location с широтой и долготой в десятичных градусах
type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func main() {
	elysium := location{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}

	bytes, err := json.MarshalIndent(elysium, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
