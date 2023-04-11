package main

import (
	"log"
	"net/http"
	"time"

	"github.com/anonIot/srvgw/handler"
	"github.com/anonIot/srvgw/repository"
	"github.com/anonIot/srvgw/services"
	"github.com/goburrow/modbus"
	"github.com/gorilla/mux"
)

func main() {

	rtuCon := initRtuConfig()

	client := repository.NewRtuBridgeDevice(rtuCon)
	rtuSev := services.NewRtuBridgeServiceDevice(client)
	acHandler := handler.NewRtuBridgeHandler(rtuSev)

	router := mux.NewRouter()

	//router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}", acHandler.GetAcIndoor).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/cmd/{slaveID:[0-9]+}/{bmsID:[0-9]+}/{cmd:[aA-zZ]+}/{val:[0-9]+}", acHandler.GetAcCmd).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/power/{val:[0-1]+}", acHandler.GetAcPower).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/temp/{val:[0-9]+}", acHandler.GetAcTemp).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/mode/{val:[0-9]+}", acHandler.GetAcMode).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/fanspeed/{val:[0-9]+}", acHandler.GetFanSpeed).Methods(http.MethodGet)
	router.HandleFunc("/api/v2/indoor/{slaveID:[0-9]+}/{bmsID:[0-9]+}/swing/{val:[0-9]+}", acHandler.GetSwing).Methods(http.MethodGet)
	// router.HandleFunc("/indoor/{slave:[0-9]+}/{bms:[0-9]+}/power/{val:[0-1]}", func(w http.ResponseWriter, r *http.Request) {

	// 	vars := mux.Vars(r)
	// 	slaveId, _ := strconv.Atoi(vars["slave"])
	// 	bms, _ := strconv.Atoi(vars["bms"])
	// 	powerVal, _ := strconv.Atoi(vars["val"])

	// 	//client := repository.NewAcRespositoryDB(rtuCon)
	// 	client := repository.NewRtuBridgeDevice(rtuCon)
	// 	rtuSev := services.NewRtuBridgeServiceDevice(client)
	// 	cmd := services.AcInddorRequest{
	// 		SlaveId: slaveId,
	// 		BmsId:   bms,
	// 		Cmd:     "power",
	// 		Value:   powerVal,
	// 	}
	// 	res, err := rtuSev.GetAcAction(cmd)
	// 	//res, err := client.AcAction(slaveId, bms, 1000, powerVal)

	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		log.Printf("Error: %v", err)
	// 		return
	// 	}

	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Fprintf(w, "%v", res)

	// }).Methods("GET")

	err := http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatalf("HTTP Server : %v", err)
	} else {

	}

}

func initRtuConfig() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler("/dev/cu.usbserial-120")
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
