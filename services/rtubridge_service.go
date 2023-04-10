package services

import (
	"fmt"
	"strings"

	"github.com/anonIot/srvgw/repository"
)

type RtuBridgeServiceDevice struct {
	RtuDevice repository.AcIndoorRepository
}

func NewRtuBridgeServiceDevice(RtuDevice repository.AcIndoorRepository) RtuBridgeService {
	return RtuBridgeServiceDevice{RtuDevice: RtuDevice}
}

func (s RtuBridgeServiceDevice) GetAcValue(slaveID int, bmsId int) (*AcIndoorInfo, error) {

	result, err := s.RtuDevice.AcRead(slaveID, bmsId)
	if err != nil {
		return nil, err
	}
	acInfo := AcIndoorInfo{
		SlaveId:   result.SlaveId,
		Bms:       result.Bms,
		Value1000: result.Value1000,
		Timer:     result.Timer,
	}

	return &acInfo, nil
}
func (s RtuBridgeServiceDevice) GetAcAction(pae AcInddorRequest) (*AcIndoorInfo, error) {

	cmd := strings.ToLower(pae.Cmd)

	switch cmd {
	case "power":
		fmt.Println(cmd)
	}

	return nil, nil

	/*

		acAddr := int(1000)
		result, err := s.RtuDevice.AcAction(pae.SlaveId, pae.BmsId, acAddr, pae.Value)
		if err != nil {
			return nil, err
		}
		acInfo := AcIndoorInfo{
			SlaveId:   result.SlaveId,
			Bms:       result.Bms,
			Value1000: result.Value1000,
			Timer:     result.Timer,
		}

		return &acInfo, nil

	*/

}
