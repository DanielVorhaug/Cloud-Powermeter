package cloud_interface

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

var SERVICE_ACCOUNT_KEY_ID string = os.Getenv("DT_SERVICE_ACCOUNT_KEY_ID")
var SERVICE_ACCOUNT_SECRET string = os.Getenv("DT_SERVICE_ACCOUNT_SECRET")
var SENSOR_ID string = os.Getenv("DT_SENSOR_ID2")
var PROJECT_ID string = os.Getenv("DT_PROJECT_ID")

var URL string = "https://emulator.d21s.com/v2/projects/" + PROJECT_ID + "/devices/" + SENSOR_ID + ":publish"

func Post_datapoint_request(channel_data <-chan float32) {
	for {
		datapoint := <-channel_data
		post_datapoint(datapoint)
	}
}

func post_datapoint(datapoint float32) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("Panic: %+v\n", x)
		}
	}()

	var data string = "{\"temperature\": {\"value\": " + fmt.Sprint(datapoint) + "}}"
	body := bytes.NewBufferString(data)

	client := http.Client{Timeout: 10 * time.Second}

	req, _ := http.NewRequest(http.MethodPost, URL, body)

	req.SetBasicAuth(SERVICE_ACCOUNT_KEY_ID, SERVICE_ACCOUNT_SECRET)
	res, _ := client.Do(req)

	res.Body.Close()
}
