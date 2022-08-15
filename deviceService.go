package device

import (
	"fmt"
)

func GetDeviceById(id string) (d Device) {
	//if devices is null call loadfile
	if len(DevicesGlobal.Devices) == 0 {
		fmt.Println("Crap file didn't load")
	}

	var device Device
	for _, dev := range DevicesGlobal.Devices {
		if dev.Id == id {
			device = dev
			break
		}

	}
	return device
	//Search devices for id
}

func getAllDevice(r *mux.Router) {
	r.HandleFunc("/device", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		d := device.GetAllDevices()
		payload, err := json.Marshal(d)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Write(payload)
		}

	})
}

