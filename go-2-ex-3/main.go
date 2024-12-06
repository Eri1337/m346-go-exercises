package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{104: "first", 117: "second", 346: "third"}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[202] = "fourth"
	// TODO: replace one
	modules[346] = "improved version of third"

	fmt.Println(modules)
}
