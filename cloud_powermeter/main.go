package main

import (
	"cloud_powermeter/cloud_interface"
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin = rpio.Pin(2)
	MESSAGE_INTERVAL int64 = 500 // [milliseconds]
	BLINKS_PER_KWH float32 = 2000.0 
)

func detect_blinks(channel_blink chan<- bool) {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()
	pin.PullUp()

	for {
		if pin.Read() == rpio.High  {
			channel_blink <- true

			for pin.Read() == rpio.High {}
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func track_blinks(channel_blink <-chan bool, channel_data chan<- float32) {

	var blink_count int64 = 0
	last_blink_time             := time.Now().UnixNano()
	sampling_period_start_time  := time.Now().UnixNano()
	next_message_time           := sampling_period_start_time + MESSAGE_INTERVAL*1000000

	for {		
		select {
		case <-channel_blink:
			blink_count++

			second_to_last_blink_time := last_blink_time
			last_blink_time = time.Now().UnixNano()
			fmt.Printf("Power: %fW", calculate_power(1, last_blink_time, second_to_last_blink_time))

			if last_blink_time > next_message_time {
				channel_data <- calculate_power(blink_count, last_blink_time, sampling_period_start_time)

				for next_message_time < time.Now().UnixNano() {
					next_message_time += MESSAGE_INTERVAL*1000000
				}
				sampling_period_start_time = last_blink_time
				blink_count = 0
			}
		}
	}
}

func calculate_power (blink_count int64, last_blink_time int64, sampling_period_start_time int64) float32 {
	var hours_measured float32 = float32(last_blink_time - sampling_period_start_time) / (1000000000.0 * 3600.0)
	var watts float32 = 1000.0 * float32(blink_count) / (BLINKS_PER_KWH * hours_measured)
	return watts
}

func main() {

	channel_blink := make(chan bool)
	channel_data := make(chan float32)
	
	go detect_blinks(channel_blink)
	go track_blinks(channel_blink, channel_data)
	go cloud_interface.Post_datapoint(channel_data)

	for {}
}
