package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"

	fmt.Println(elements["He"])

	if name, ok := elements["foo"]; ok {
		fmt.Println(name)
	}

	if name, ok := elements["H"]; ok {
		fmt.Println(name)
	}

	elements = map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	for k, v := range elements {
		fmt.Println(k, "->", v)
	}

	if el, ok := elements["Be"]; ok {
		fmt.Println(el)
	}
}
