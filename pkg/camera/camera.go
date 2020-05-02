package camera

import (
	"fmt"
	"time"
)

type Camera struct {
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) ShootImage() (string, error) {
	return fmt.Sprintf("%s.jpeg", time.Now().Format("2006-01-02T15:04:05")), nil
}
