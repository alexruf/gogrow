package sensor

type TemperaturSensor struct {
}

func NewTemperaturSensor() *TemperaturSensor {
	return &TemperaturSensor{}
}

func (c *TemperaturSensor) Query() (int32, error) {
	return 0, nil
}
