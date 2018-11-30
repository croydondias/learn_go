package main

import (
	"fmt"
)

var pow = []int{1, 5, 6, 3, 10}

type Vertex struct {
	Lat  float64
	Long float64
}

func main() {
	fmt.Println(len(pow), cap(pow))
	for i, k := range pow {
		fmt.Println("index = ", i, " value =", k)

	}

	var m = make(map[string]Vertex)
	m["Croydon"] = Vertex{12.3245, -33.234}
	fmt.Println(m["Croydon"])
}
