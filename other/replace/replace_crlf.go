package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	absPath, _ := filepath.Abs("test.txt")
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Println(input)

	re := regexp.MustCompile("  ")
	re2 := regexp.MustCompile(" ")
	input = re.ReplaceAllString(input, "\r\n")
	input = re2.ReplaceAllString(input, "\n")
	fmt.Println(input)

	f, _ := os.Create("data.txt")
	_, err2 := f.WriteString(input)
	if err2 != nil {
		panic(err)
	}

}
