package main

import (
	"fmt"
	"runtime"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(">> ", i)

	}
	fmt.Println(sum)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("Other")
	}
}
