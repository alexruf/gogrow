package sensor

type HumiditySensor struct {
}

func NewHumiditySensor() *HumiditySensor {
	return &HumiditySensor{}
}

func (c *HumiditySensor) Query() (int32, error) {
	return 0, nil
}
