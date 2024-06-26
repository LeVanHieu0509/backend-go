package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// peace := "peace"
	// var peace = "peace"
	// var peace string = "peace"

	// var blank string

	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)

	//Multiple-line raw string literals
	fmt.Println(`peace be upon you upon
	 you be peace`)

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)

	type byte = uint8
	type rune = int32

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	//To display the characters rather than their numeric values, the %c format
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	// grade := 'A'
	// var star byte = '*'

	//A variable can be assigned to a different string, but strings themselves can’t be altered:
	// peace := "shalom"
	// peace = "salām"

	//index string
	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c)

	/*
			QC 9.2 answer
		1 128 characters.
		2 A byte is an alias for the uint8 type. A rune is an alias for the int32 type.
		3 var star byte = '*' fmt.Printf("%c %[1]v\n", star)
		smile := ''
		fmt.Printf("%c %[1]v\n", smile)
		   acute := 'é'
		   fmt.Printf("%c %[1]v\n", acute)
	*/

	var a uint = 128
	fmt.Println(a)

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)
	smile := ''
	fmt.Printf("%c %[1]v\n", smile)
	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)

	//
	//Manipulating characters with Caesar cipher
	ci := 'a'
	ci = ci + 3
	fmt.Printf("%c\n", ci)

	//Manipulate a single character
	message1 := "shalom"
	for i := 0; i < 6; i++ {
		c := message1[i]
		fmt.Printf("%c\n", c)
	}

	message2 := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message2); i++ {
		c := message2[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

	//UTF8
	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	// c, size := utf8.DecodeRuneInString(question)
	// fmt.Printf("First rune: %c %v bytes", c, size)

	//Decoding runes
	question1 := "¿Cómo estás?"
	for i, c := range question1 {
		fmt.Printf("%v %c\n", i, c)
	}

	/*
	   How many runes are in the English alphabet "abcdefghijklmnopqrstuvwxyz"? How many bytes?
	  2 How many bytes are in the rune '¿'?


	*/

	/*
		1 There are 26 runes and 26 bytes in the English alphabet.
		2 There are 2 bytes in the rune '¿'.
	*/
}
