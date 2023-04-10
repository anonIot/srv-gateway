package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anonIot/srvgw/services"
)

type rtuBridgeHandler struct {
	rtuSrv services.RtuBridgeService
}

func NewRtuBridgeHandler(rtuSrv services.RtuBridgeService) rtuBridgeHandler {

	return rtuBridgeHandler{rtuSrv: rtuSrv}
}

func (h rtuBridgeHandler) GetAcIndoor(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This Get Indoor")
	acinfo, err := h.rtuSrv.GetAcValue(1, 1)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(acinfo)

}
