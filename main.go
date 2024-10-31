package main

import (
	"wepa.com/go/mserial"
	"wepa.com/go/mserver"
)



func main() {
	mserial.CreateConnection()
	mserver.CreateApi()

}
