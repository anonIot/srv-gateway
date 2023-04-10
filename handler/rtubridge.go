package handler

import (
	"encoding/json"
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

	//vars := mux.Vars(r)
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

func (h rtuBridgeHandler) GetAcAction(w http.ResponseWriter, r *http.Request) {}
