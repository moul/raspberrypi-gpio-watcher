package main

import "fmt"
import "github.com/stianeikeland/go-rpio"

func main() {
	fmt.Println("Initializing...")
	err := rpio.Open()
	fmt.Println(err)
}
