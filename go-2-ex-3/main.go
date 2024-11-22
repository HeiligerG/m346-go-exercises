package main

import "fmt"

func main() {
	modules := map[int]int16{
		104: 10,
		117: 20,
		346: 30,
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)

	// TODO: add one
	modules[500] = 40

	// TODO: replace one
	modules[346] = 35

	fmt.Println(modules)
}
