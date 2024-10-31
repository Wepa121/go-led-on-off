package mserial

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

var port serial.Port

func CreateConnection() {
	mode := &serial.Mode {
		BaudRate: 9600,
	}
	
	port, _ = serial.Open("/dev/ttyUSB0", mode)
	time.Sleep(time.Second * 2)
}

func SendChar(c string) {
	n, err := port.Write([]byte(fmt.Sprintf("%v\n\r", c)))

	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Sent %v bytes \n", n)
}
