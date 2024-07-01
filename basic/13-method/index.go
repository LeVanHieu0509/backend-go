package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type celsius float64
type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	return 0
}

func main() {

	var temperature celsius = 20
	fmt.Printf("%T - %[1]v \n", temperature)

	const degrees = 20
	temperature += 10

	// temperature += warmUp //invalid

	var warmUp float64 = 10

	//convert to same type
	temperature += celsius(warmUp)

	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "o K is ", c, "o C")

	var c1 celsius = 127.0
	k1 := celsiusToKelvin(c1)
	fmt.Println(c, "o C is ", k1, "o K")

	fmt.Println("-------------FUNCTION FIRST -------------------")

	// First- function
	sensor := fakeSensor
	fmt.Println("sensor 1:", sensor())
	sensor = realSensor
	fmt.Println("sensor 2:", sensor())

	measureTemperature(3, fakeSensor)

}

// kelvinToCelsius converts oK to oC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// ép kiểu thôi
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

/*

	In Go you can assign functions to variables, pass functions to functions, and even write functions that return functions

*/

// First function
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}
