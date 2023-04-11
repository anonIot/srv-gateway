package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anonIot/srvgw/services"
	"github.com/gorilla/mux"
)

type rtuBridgeHandler struct {
	rtuSrv services.RtuBridgeService
}

func NewRtuBridgeHandler(rtuSrv services.RtuBridgeService) rtuBridgeHandler {

	return rtuBridgeHandler{rtuSrv: rtuSrv}
}

func (h rtuBridgeHandler) GetAcIndoor(w http.ResponseWriter, r *http.Request) {

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])

	acinfo, err := h.rtuSrv.GetAcValue(slaveID, bmsID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(acinfo)

}

func (h rtuBridgeHandler) GetAcCmd(w http.ResponseWriter, r *http.Request) {

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])
	cmd, _ := mux.Vars(r)["cmd"]
	value, _ := strconv.Atoi(mux.Vars(r)["val"])

	switch cmd {
	case "power":

		fmt.Println(cmd)
		addr := (1000 + (bmsID * 10) - 1)
		fmt.Println(addr)
		cmdRequest := services.AcInddorRequest{
			SlaveId: slaveID,
			BmsId:   bmsID,
			Addr:    addr,
			Value:   value,
		}

		result, err := h.rtuSrv.GetAcAction(cmdRequest)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)

	case "temp":

		temp := value * 2
		addr := (1000 + (bmsID * 10) - 1)
		addr = addr + 2

		fmt.Println(addr)

		cmdRequest := services.AcInddorRequest{
			SlaveId: slaveID,
			BmsId:   bmsID,
			Addr:    addr,
			Value:   temp,
		}

		result, err := h.rtuSrv.GetAcAction(cmdRequest)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		fmt.Println(bmsID)

	case "fan":
		fmt.Println(cmd)
	case "speed":
		fmt.Println(cmd)
	case "mode":
		fmt.Println(cmd)
	case "swing":
		fmt.Println(cmd)

	}
	// 	cmd := services.AcInddorRequest{
	// 		SlaveId: slaveId,
	// 		BmsId:   bms,
	// 		Cmd:     "power",
	// 		Value:   powerVal,
	// 	}

	// result, err := h.rtuSrv.GetAcAction()
	// if err != nil {
	// 	handleError(w, err)
	// 	return
	// }

	// fmt.Println(result)
	// w.Header().Set("content-type", "application/json")
	// json.NewEncoder(w).Encode(result)
}

func (h rtuBridgeHandler) GetAcPower(w http.ResponseWriter, r *http.Request) {

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])
	value, _ := strconv.Atoi(mux.Vars(r)["val"])

	addr := 1000 + (bmsID * 10) - 1
	cmdRequest := services.AcInddorRequest{
		SlaveId: slaveID,
		BmsId:   bmsID,
		Addr:    addr,
		Value:   value,
	}

	result, err := h.rtuSrv.GetAcAction(cmdRequest)

	if err != nil {
		handleError(w, err)
		return
	}
	acinfo, err := h.rtuSrv.GetAcValue(slaveID, bmsID)
	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Println(result)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(acinfo)

}
func (h rtuBridgeHandler) GetAcTemp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Temp")

	//	sampleRegex := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])
	//value, _ := strconv.Atoi(mux.Vars(r)["val"])
	val := mux.Vars(r)["val"]
	value, _ := strconv.ParseFloat(val, 64)
	temp := int(value * 2)

	// fmt.Printf("temp %v", val)
	// fmt.Printf("temp %v", temp)

	// return

	if temp >= 30 && value <= 60 {

		addr := (1000 + (bmsID * 10) - 1)
		addr = addr + 2
		fmt.Println(addr)
		cmdRequest := services.AcInddorRequest{
			SlaveId: slaveID,
			BmsId:   bmsID,
			Addr:    addr,
			Value:   temp,
		}

		result, err := h.rtuSrv.GetAcAction(cmdRequest)

		if err != nil {
			handleError(w, err)
			return
		}
		acinfo, err := h.rtuSrv.GetAcValue(slaveID, bmsID)
		if err != nil {
			handleError(w, err)
			return
		}

		fmt.Println(result)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(acinfo)

	} else {
		errsg := errors.New("Temp is value rang : 15-30")
		handleError(w, errsg)
		return
	}

}

func (h rtuBridgeHandler) GetAcMode(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Temp")

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])
	value, _ := strconv.Atoi(mux.Vars(r)["val"])

	if value >= 0 && value <= 6 {
		mode := value
		addr := (1000 + (bmsID * 10) - 1)
		addr = addr + 1
		fmt.Println(addr)
		cmdRequest := services.AcInddorRequest{
			SlaveId: slaveID,
			BmsId:   bmsID,
			Addr:    addr,
			Value:   mode,
		}

		result, err := h.rtuSrv.GetAcAction(cmdRequest)

		if err != nil {
			handleError(w, err)
			return
		}
		acinfo, err := h.rtuSrv.GetAcValue(slaveID, bmsID)
		if err != nil {
			handleError(w, err)
			return
		}

		fmt.Println(result)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(acinfo)

	} else {
		errsg := errors.New("Mode is value rang : 15-30")
		handleError(w, errsg)
		return
	}

}

func (h rtuBridgeHandler) GetFanSpeed(w http.ResponseWriter, r *http.Request) {

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])
	value, _ := strconv.Atoi(mux.Vars(r)["val"])

	if value >= 0 && value <= 4 {
		speed := value
		addr := (1000 + (bmsID * 10) - 1)
		addr = addr + 6
		fmt.Println(addr)
		cmdRequest := services.AcInddorRequest{
			SlaveId: slaveID,
			BmsId:   bmsID,
			Addr:    addr,
			Value:   speed,
		}

		result, err := h.rtuSrv.GetAcAction(cmdRequest)

		if err != nil {
			handleError(w, err)
			return
		}
		acinfo, err := h.rtuSrv.GetAcValue(slaveID, bmsID)
		if err != nil {
			handleError(w, err)
			return
		}

		fmt.Println(result)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(acinfo)

	} else {
		errsg := errors.New("Fan Speed is Rang : 0-4")
		handleError(w, errsg)
		return
	}
}

func (h rtuBridgeHandler) GetSwing(w http.ResponseWriter, r *http.Request) {

	slaveID, _ := strconv.Atoi(mux.Vars(r)["slaveID"])
	bmsID, _ := strconv.Atoi(mux.Vars(r)["bmsID"])
	value, _ := strconv.Atoi(mux.Vars(r)["val"])

	if value >= 0 && value <= 6 {
		speed := value
		addr := (1000 + (bmsID * 10) - 1)
		addr = addr + 7
		fmt.Println(addr)
		cmdRequest := services.AcInddorRequest{
			SlaveId: slaveID,
			BmsId:   bmsID,
			Addr:    addr,
			Value:   speed,
		}

		result, err := h.rtuSrv.GetAcAction(cmdRequest)

		if err != nil {
			handleError(w, err)
			return
		}
		acinfo, err := h.rtuSrv.GetAcValue(slaveID, bmsID)
		if err != nil {
			handleError(w, err)
			return
		}

		fmt.Println(result)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(acinfo)

	} else {
		errsg := errors.New("Louver is value rang : 0-5")
		handleError(w, errsg)
		return
	}
}
