# GoGrow

GoGrow is an open-source project using a Raspberry Pi with sensors and camera to watch and monitor your plants grow.

> Note: This project is currently in a very early development stage and might not be ready for general usage!

## About GoGrow

* At the beginning, the target is to query the temperature and humidity at regular intervals from the sensor connected to a Raspberry Pi and generate a camera image with the respective values
* The camera images should be used to create a timelapse video where you can watch your plants grow
* Sensor data can be used to generate diagrams and evaluations
* Support of additional hardware and sensors
* Add custom rules and conditions to trigger scripts to allow automated workflows

## Technology

This project is mainly being written in [Golang](https://golang.org/).

## Supported hardware

Currently, only a very limited set and hardware is supported.
Support of additional hardware and sensors may follow later.

### Boards

* [Raspberry Pi 4 Model B](https://www.raspberrypi.org/products/raspberry-pi-4-model-b/)
* [Raspberry Pi 3 Model B+](https://www.raspberrypi.org/products/raspberry-pi-3-model-b-plus/)

### Sensors

* [Raspberry Pi Camera Module V2](https://www.raspberrypi.org/products/camera-module-v2/)
* [Raspberry Pi Sense HAT](https://www.raspberrypi.org/products/sense-hat/)
