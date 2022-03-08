package main

import (
	"cloud_powermeter/cloud_interface"
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
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

var (
	pin = rpio.Pin(2)
	BLINK_INTERVAL int64 = 2 // [seconds]
)

func track_blinks(channel_blink chan<- bool) {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()
	pin.PullDown()
	//pin.Detect(rpio.FallEdge) // enable falling edge event detection

	//defer pin.Detect(rpio.NoEdge)

	for {
		// state := pin.Read() 
		// fmt.Println(state)
		if pin.Read()==rpio.High  {//pin.EdgeDetected() { // check if event occured
			fmt.Println("Blink detected!")
			channel_blink <- true
			for pin.Read()==rpio.High {
				time.Sleep(1 * time.Millisecond)
			}
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func track_blinks_dummy(channel_blink chan<- bool) {

	start_time  := time.Now().UnixNano()
	for {
		time.Sleep(200*time.Millisecond)
		fmt.Printf("\nBlink detected at %f\n", (time.Now().UnixNano()-start_time))
		channel_blink <- true
	}
}

func calculate_power(channel_blink <-chan bool, channel_data chan<- float32) {

	var blink_count int64 = 0
	last_blink_time             := time.Now().UnixNano()
	sampling_period_start_time  := time.Now().UnixNano()
	next_message_time           := sampling_period_start_time + BLINK_INTERVAL*1000000000

	for {		
		select {
		case <-channel_blink:
			blink_count++
			fmt.Printf("\nBlink number is %d\n", blink_count)
			last_blink_time = time.Now().UnixNano()

			if last_blink_time > next_message_time {
				var data float32 = float32(blink_count * 1000000000) / float32(last_blink_time - sampling_period_start_time) // Blinks per second in period
				sampling_period_start_time = last_blink_time
				next_message_time = sampling_period_start_time + BLINK_INTERVAL*1000000000
				// fmt.Printf("\nseconds passed, transmitting %f", data)
				channel_data <- data
				blink_count = 0
			}
		}
	}
}

func main() {
	fmt.Println("Started...")

	
	channel_blink := make(chan bool)
	channel_data := make(chan float32)
	
	go track_blinks(channel_blink)
	//go track_blinks_dummy(channel_blink)
	go calculate_power(channel_blink, channel_data)
	go cloud_interface.Post_datapoint(channel_data)

	fmt.Println("Ready!")

	for {}

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
	//return
}
