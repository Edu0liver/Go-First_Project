package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

var pl = fmt.Println
var tp = reflect.TypeOf

func main() {
	casting()
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

func casting() {
	cV1 := 1.5
	cV2 := int(cV1)

	pl(cV2)
	pl(tp(cV2))

	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)

	pl(cV4, err, tp(cV4))

	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)

	pl(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8, err, tp(cV8))
	} else {
		log.Fatal(err)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9, tp(cV9))
}

func logicalConditions() {
	// Conditional Operators ==, !=, >, <, >=, <=
	// Logical Operators: &&, ||, !

	iAge := 18

	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Bithday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Bithday")
	} else if iAge >= 65 {
		pl("Not Important Bithday")
	} else {
		pl("Not Important Bithday")
	}

	pl("!true", !true)

}
