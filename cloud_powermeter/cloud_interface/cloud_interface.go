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
var SENSOR_ID string = os.Getenv("DT_SENSOR_ID")
var PROJECT_ID string = os.Getenv("DT_PROJECT_ID")

var URL string = "https://emulator.d21s.com/v2/projects/" + PROJECT_ID + "/devices/" + SENSOR_ID + ":publish"

func Post_datapoint(channel_data <-chan float32) {
	for {
		select {
		case datapoint := <-channel_data:
			timeBegin := time.Now().UnixNano()
			var data string = "{\"temperature\": {\"value\": " + fmt.Sprint(datapoint) + "}}"
			body := bytes.NewBufferString(data)

			client := http.Client{Timeout: 10 * time.Second}

			req, _ := http.NewRequest(http.MethodPost, URL, body)
			req.SetBasicAuth(SERVICE_ACCOUNT_KEY_ID, SERVICE_ACCOUNT_SECRET)
			res, _ := client.Do(req)

			defer res.Body.Close()
			// resBody, _ := io.ReadAll(res.Body)
			// fmt.Printf("Status: %d\n", res.StatusCode)
			// fmt.Printf("Body: %s\n", string(resBody))
			fmt.Printf("Time taken: %f\n", float32(time.Now().UnixNano() - timeBegin)/1000000000.0)
		}
	}
}

func Test() {
	fmt.Println(URL)
	fmt.Println(SERVICE_ACCOUNT_KEY_ID)
	fmt.Println(SERVICE_ACCOUNT_SECRET)
	fmt.Println(SENSOR_ID)
	fmt.Println(PROJECT_ID)

}
