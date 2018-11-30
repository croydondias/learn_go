package main

import "fmt"
import "log"

func add(x int, y int) int {
	return x + y
}

func main() {
	var name string = "Croydon"
	fmt.Println(name)
	fmt.Println(add(10, 50))

	who := "Miaw"
	log.Print(who)
}

