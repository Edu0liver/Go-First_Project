package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println
var tp = reflect.TypeOf
var sc = fmt.Scan

func main() {
	loops()
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

	iAge := 0

	sc(&iAge)

	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Bithday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Bithday")
	} else if iAge >= 65 {
		pl("Important Bithday")
	} else {
		pl("Not Important Bithday")
	}

	pl("!true =", !true)

}

func stringsManipulate() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)

	pl(sV2)
	pl("Length:", len(sV2))
	pl("Contains Another:", strings.Contains(sV2, "Another"))
	pl("'o' index:", strings.Index(sV2, "o"))
	pl("Replace:", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome words\n"
	sV3 = strings.TrimSpace(sV3)

	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("Lower:", strings.ToLower(sV2))
	pl("Upper:", strings.ToUpper(sV2))
	pl("Prefix:", strings.HasPrefix("tococat", "toco"))
	pl("Suffix:", strings.HasSuffix("tococat", "cat"))
}

func runes() {
	rStr := "abcdefg"

	pl("Rune Count:", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}

func timeManipulate() {
	now := time.Now()

	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

}

func mathMain() {
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))

	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)

	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)

	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guessesbased on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)

	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)

}

func loops() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for true {

		fmt.Print("Guess a number between 0 and 50: ")
		pl("Random Number is:", randNum)

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)

		if iGuess > randNum {
			pl("Pick a Lower Value!")
		} else if iGuess < randNum {
			pl("Pick a Higher Value!")
		} else {
			pl("You Guessed it!")
			break
		}
	}
}
