package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anonIot/srvgw/repository"
	"github.com/goburrow/modbus"
)

func main() {

	rtuCon := initRtuConfig()
	client := repository.NewAcRespositoryDB(rtuCon)
	results, err := client.AcAction(1, 1)

	if err != nil {
		log.Fatalf("Error Read %v", err)
		return
	}

	fmt.Println(results)

}

func initRtuConfig() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler("/dev/cu.usbserial-1120")
	handler.BaudRate = 19200
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 3 * time.Second

	err := handler.Connect()
	defer handler.Close()

	if err != nil {
		log.Fatalf(" No Connect : %v", err)
		return nil
	}
	return handler
}
