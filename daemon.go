package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/stianeikeland/go-rpio"
)

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
		if pin.Read() == 1 {
			cmd := exec.Command("/etc/gpio-trigger")
			err := cmd.Start()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Waiting for command to finish...\n")
			err = cmd.Wait()
			fmt.Printf("Command finished with error: %v\n", err)
			os.Exit(1)
		}
		time.Sleep(time.Second / 5)
	}
}
