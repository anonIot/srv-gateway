package repository

type AcPacketRepository struct {
	SlaveId   int
	Bms       int
	Value1000 []byte
	Timer     string
}

type AcIndoorRepository interface {
	AcScan() ([]AcPacketRepository, error)
	AcAction(int, int, int) (*AcPacketRepository, error)
	AcRead(int, int) (*AcPacketRepository, error)
}
