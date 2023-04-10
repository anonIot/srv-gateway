package repository

import (
	"time"

	"github.com/goburrow/modbus"
)

type rtuBridgeDevice struct {
	Cli *modbus.RTUClientHandler
}

func NewRtuBridgeDevice(client *modbus.RTUClientHandler) AcIndoorRepository {
	return rtuBridgeDevice{Cli: client}
}

func (r rtuBridgeDevice) AcScan() ([]AcScanRepository, error) {
	return nil, nil
}

func (r rtuBridgeDevice) AcRead(slaveId int, bmsId int) (*AcPacketRepository, error) {
	sid := slaveId
	bms := bmsId

	handler := r.Cli
	handler.SlaveId = byte(sid)
	acAddress := (1000 + (bms * 10) - 1)
	client := modbus.NewClient(handler)
	result, err := client.ReadHoldingRegisters(uint16(acAddress), uint16(10))

	if err != nil {
		return nil, err
	}
	now := time.Now()

	acInfo := AcPacketRepository{
		SlaveId:   sid,
		Bms:       bmsId,
		Value1000: result,
		Timer:     now.String(),
	}

	return &acInfo, nil
}

func (r rtuBridgeDevice) AcAction(slaveID int, bmsID int, addr int, val int) (*AcPacketRepository, error) {
	slaveId := slaveID
	bms := bmsID

	acAddress := addr

	handler := r.Cli
	handler.SlaveId = byte(slaveId)
	client := modbus.NewClient(handler)
	result, err := client.WriteSingleRegister(uint16(acAddress), uint16(val))

	if err != nil {
		return nil, err
	}
	now := time.Now()

	acInfo := AcPacketRepository{
		SlaveId:   slaveId,
		Bms:       bms,
		Value1000: result,
		Timer:     now.String(),
	}

	return &acInfo, nil
}
