package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	variables()
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
