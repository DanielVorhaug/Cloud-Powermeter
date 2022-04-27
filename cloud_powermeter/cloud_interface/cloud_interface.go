package cloud_interface

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var SERVICE_ACCOUNT_KEY_ID string = os.Getenv("DT_SERVICE_ACCOUNT_KEY_ID")
var SERVICE_ACCOUNT_SECRET string = os.Getenv("DT_SERVICE_ACCOUNT_SECRET")
var SENSOR_ID string = os.Getenv("DT_SENSOR_ID")
var PROJECT_ID string = os.Getenv("DT_PROJECT_ID")

var URL string = "https://emulator.d21s.com/v2/projects/" + PROJECT_ID + "/devices/" + SENSOR_ID + ":publish"

func Post_datapoint(channel_data <-chan float32) {
	for {
		datapoint := <-channel_data

		var data string = "{\"temperature\": {\"value\": " + fmt.Sprint(datapoint) + "}}"
		body := bytes.NewBufferString(data)

		client := http.Client{Timeout: 10 * time.Second}

		req, err := http.NewRequest(http.MethodPost, URL, body)
		if err != nil {
			log.Fatalln(err)
		}

		req.SetBasicAuth(SERVICE_ACCOUNT_KEY_ID, SERVICE_ACCOUNT_SECRET)
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		res.Body.Close()
	}
}
