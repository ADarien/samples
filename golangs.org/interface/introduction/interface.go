package main

import (
	"fmt"
	"strings"
)

// talk
var t interface {
	Talk() string
}

type martian struct{}

// martian talk
func (m martian) Talk() string {
	return "nack nack"
}

type laser int

// laser talk
func (l laser) Talk() string {
	return strings.Repeat("pew ", int(l))
}

// тот, кто говорит
type talker interface {
	Talk() string
}

// выстрел с параметром talker
func shout(t talker) {
	louder := strings.ToUpper(t.Talk())
	fmt.Println(louder)
}

type rover string

// rover talk
func (r rover) Talk() string {
	return string(r)
}

// встраиваем laser в starship
type starship struct {
	laser
}

func main() {
	t = martian{}
	fmt.Println(t.Talk())

	t = laser(3)
	fmt.Println(t.Talk())

	shout(martian{})
	shout(laser(2))

	r := rover("whir whir")
	shout(r)

	//когда говорит starship, за него говорит laser
	s := starship{laser(3)}
	fmt.Println(s.Talk())
	shout(s)
}
