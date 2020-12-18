package main

import (
	"fmt"
	"strings"
)

// talk
var t interface {
	talk() string
}

type martian struct{}

// martian talk
func (m martian) talk() string {
	return "nack nack"
}

type laser int

// laser talk
func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// тот, кто говорит
type talker interface {
	talk() string
}

// выстрел с параметром talker
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type rover string

// rover talk
func (r rover) talk() string {
	return string(r)
}

// встраиваем laser в starship
type starship struct {
	laser
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))

	r := rover("whir whir")
	shout(r)

	//когда говорит starship, за него говорит laser
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}
