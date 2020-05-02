package snapshot

import (
	"fmt"
	"github.com/alexruf/gogrow/pkg/camera"
	"github.com/alexruf/gogrow/pkg/sensor"
)

func CaptureSnapshot() error {
	cam := camera.NewCamera()
	if i, err := cam.ShootImage(); err != nil {
		return err
	} else {
		fmt.Printf("Image: %s\n", i)
	}

	temp := sensor.NewTemperaturSensor()
	if t, err := temp.Query(); err != nil {
		return err
	} else {
		fmt.Printf("Temperatur: %d\n", t)
	}

	humi := sensor.NewHumiditySensor()
	if h, err := humi.Query(); err != nil {
		return err
	} else {
		fmt.Printf("Humidity: %d\n", h)
	}

	return nil
}
