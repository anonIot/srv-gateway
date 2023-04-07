package services

import "github.com/anonIot/srvgw/repository"

type RtuBridgeServiceDevice struct {
	RtuDevice repository.AcIndoorRepository
}

func NewRtuBridgeServiceDevice(RtuDevice repository.AcIndoorRepository) RtuBridgeService {
	return RtuBridgeServiceDevice{RtuDevice: RtuDevice}
}

func (s RtuBridgeServiceDevice) GetAcValue(slaveID int, bmsId int) (*AcIndoorInfo, error) {

	return nil, nil
}
func (s RtuBridgeServiceDevice) GetAcAction(pae AcInddorRequest) (*AcInddorRequest, error) {
	return nil, nil
}
