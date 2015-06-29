package main

import "fmt"
import "github.com/stianeikeland/go-rpio"
import "time"

func main() {
	fmt.Println("Initializing...")
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()	
	fmt.Println("Initialized !")

	pin := rpio.Pin(10)
	pin.Input()

	for {
		fmt.Printf("Result: %v\n", pin.Read())
		time.Sleep(time.Second / 5)
	}
}
