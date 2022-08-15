package device

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Device struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Organization  string `json:"organization"`
	Created_at    string `json:"created_at"`
	Serial_number string `json:"serial_number"`
	Run_time      string `json:"run_time"`
	Creator       string `json:"creator"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}

var DevicesGlobal Devices

func init() {
	//Load file and add device list to devices
	jsonFile, err := os.Open("device.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	//var devices Devices
	json.Unmarshal(byteValue, &DevicesGlobal)
	//fmt.Println(len(devices.Devices))

	jsonFile.Close()
}
