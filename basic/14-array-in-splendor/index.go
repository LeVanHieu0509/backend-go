package main

import (
	"fmt"
	"sort"
)

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)

	fmt.Println(planets[3] == "")

	fmt.Println("-----------------Initialize arrays with composite literals------------------")
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	planets1 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(dwarfs, planets1)

	fmt.Println("-----------------Looping through an array------------------")
	dwarfs1 := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	// Looping through an array cach 2
	dwarfs2 := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	fmt.Println(dwarfs1, dwarfs2)

	fmt.Println("-----------------Arrays are copied------------------")
	planetsMarkII := planets

	planets[2] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	fmt.Println("-----------------Arrays of arrays------------------")
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Printf("board - %v \n", board)

	fmt.Println("-----------------Slicing an array------------------")
	terrestrial := planets1[0:4]
	gasGiants := planets1[4:6]
	iceGiants := planets1[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(gasGiants[0])

	terrestrial1 := planets1[:4]
	gasGiants1 := planets1[4:6]
	iceGiants1 := planets1[6:]
	fmt.Println("terrestrial1:", terrestrial1, "gasGiants1:", gasGiants1, "iceGiants1:", iceGiants1)

	fmt.Println("-----------------Composite literals for slices------------------")
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	colonized := terrestrial[2:]

	fmt.Println("dwarfSlice:", dwarfSlice, "colonized:", colonized)

	fmt.Println("-----------------Slices with methods------------------")
	planets2 := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	sort.StringSlice(planets2).Sort()
	//        Sorts planets alphabetically
	fmt.Println(planets2)

}
