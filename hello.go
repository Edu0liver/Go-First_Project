package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

var pl = fmt.Println

func main() {
	dataTypes()
}

func nameQuestion() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}

func variables() {
	var vName string = "Edu"
	v1, v2 := 1.2, 3.4
	v3 := "hello"
	v4 := 2.4

	pl(vName, v1, v2, v3, v4)
}

func dataTypes() {
	// int, float, bool, string, rune
	// Default type 0, 0.0, false, ""

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
}
