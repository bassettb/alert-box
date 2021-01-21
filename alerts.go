package main

import (
	"os"
    "fmt"
    "net/http"

    "github.com/stianeikeland/go-rpio"
)

const relay1Pin int = 17

func main() {
    fmt.Println("opening gpio")
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Printf("unable to open gpio: %s\n", err.Error())
		os.Exit(-1)
    }

    defer rpio.Close()

    pin := rpio.Pin(relay1Pin)
	pin.Output()

	http.HandleFunc("/alerts/on", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Relay ON")
        pin.Write(rpio.High)
	})
	http.HandleFunc("/alerts/off", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Relay off")
        pin.Write(rpio.Low)
	})
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic("http Listen failed")
	}
}

/*
func setupButtonPin(number int) {
    pin := rpio.Pin(number)

	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge) // enable falling edge event detection

	fmt.Println("press a button")

	if pin.EdgeDetected() {}
	pin.Detect(rpio.NoEdge) // disable edge event detection
}
*/
