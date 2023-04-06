package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/goburrow/modbus"
)

type acRepositoryDB struct {
	Cli *modbus.RTUClientHandler
}

func NewAcRespositoryDB(Client *modbus.RTUClientHandler) acRepositoryDB {

	return acRepositoryDB{Cli: Client}
}

func (r acRepositoryDB) AcAction(slaveID int, bmsId int, addr int, val int) (*AcPacketRepository, error) {
	sid := slaveID
	bms := bmsId

	handler := r.Cli
	handler.SlaveId = byte(sid)
	acAddress := (addr + (bms * 10)) - 1
	client := modbus.NewClient(handler)

	// results, err := client.ReadHoldingRegisters(uint16(acAddress), uint16(10))

	// if err != nil {
	// 	return nil, err
	// }

	cmd, err := client.WriteSingleRegister(uint16(acAddress), uint16(val))

	if err != nil {
		log.Fatalf("cmd err %v", err)
		return nil, err
	}
	fmt.Println(cmd)
	//log.Fatalf(string(cmd))

	results, err := client.ReadHoldingRegisters(uint16(acAddress), uint16(10))

	if err != nil {
		return nil, err
	}

	now := time.Now()

	acInfo := AcPacketRepository{
		SlaveId:   sid,
		Bms:       bmsId,
		Value1000: results,
		Timer:     now.String(),
	}

	return &acInfo, nil
}
func (r acRepositoryDB) AcRead(slaveID int, bmsId int) (*AcPacketRepository, error) {
	sid := slaveID
	bms := bmsId

	handler := r.Cli
	handler.SlaveId = byte(sid)
	acAddress := (1000 + (bms * 10)) - 1
	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(uint16(acAddress), uint16(10))

	if err != nil {
		return nil, err
	}
	now := time.Now()

	acInfo := AcPacketRepository{
		SlaveId:   sid,
		Bms:       bmsId,
		Value1000: results,
		Timer:     now.String(),
	}

	return &acInfo, nil
}
