package main

import (
	"cloud_powermeter/cloud_interface"
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin = rpio.Pin(2)
)

// func blink_counter() {
// 	for {
// 		select {

// 		}
// 	}
// }

// func send_data(channel_data <-chan float32) {
// 	data := <- channel_data
// 	cloud_interface.Post_datapoint(data)
// 	fmt.Println("Sent")
// }

func main() {
	fmt.Println("Running...")

	channel_data := make(chan float32)

	go cloud_interface.Post_datapoint(channel_data)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	var datapoint float32 = 11.42

	// pin := rpio.Pin(15)
	// pin.Input()
	// pin.Detect(rpio.RiseEdge)

	// var blink_count int 		= 0
	// last_blink_time             = time.Now().UnixNano()
	// sampling_period_start_time  = time.Now().UnixNano()
	// next_message_time           = time.Now().UnixNano() + 5*1000000000

	// // for ;; {

	// // }
	// var a time.Time = time.Now()
	// fmt.Println(a.UnixNano())
	// time.Sleep(1 * time.Second)
	// fmt.Println(time.Now().UnixNano()-a.UnixNano())

	// // Set pin to output mode
	// pin.Output()

	// // Toggle pin 20 times
	// for x := 0; x < 20; x++ {
	// 	pin.Toggle()
	// 	time.Sleep(time.Second / 5)
	// }
	return
}
