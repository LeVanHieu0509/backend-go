package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)

	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	//            Length 4, capacity 4
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)

	dwarfs1 := make([]string, 0, 10)
	dwarfs1 = append(dwarfs1, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	fmt.Println(dwarfs1)
}
