package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// What in the current temperature in C ? 0
	// output: 0 C = 32 F
	// Would you like to convert another temperature ? (y/n) y
	// output: 40 C = 104 F
	// Would you like to convert another temperature ? (y/n) n
	// Good bye!
	degree := os.Args[1]

	for {

		fmt.Printf("What in the current temperature in %s? ", degree)

		var value1 string
		fmt.Scanln(&value1)
		value, err := strconv.ParseFloat(value1, 10)

		c, u, err := Convert(degree, value)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strconv.FormatFloat(value, 'f', -1, 64) + " " + degree + " = " + strconv.FormatFloat(c, 'f', -1, 64) + " " + u)

		fmt.Printf("Would you like to convert another temperature ? (y/n) ")

		var oper string

		fmt.Scanf("%s", &oper)
		if oper == "n" {
			fmt.Printf("Good bye!")
			break
		}
	}
}

func Convert(u string, v float64) (float64, string, error) {
	if u == "C" {
		return v*9/5 + 32, "F", nil
	} else if u == "F" {
		return (v - 32) * 5 / 9, "C", nil
	}

	return 0, u, fmt.Errorf("Error: Not a valid input")
}
