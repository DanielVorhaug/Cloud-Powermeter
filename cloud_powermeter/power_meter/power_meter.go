package power_meter

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin = rpio.Pin(2)
)

func track_blinks() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge) // enable falling edge event detection

	defer pin.Detect(rpio.NoEdge)

	for {
		if pin.EdgeDetected() { // check if event occured
			fmt.Println("button pressed")
		}
	}
}
