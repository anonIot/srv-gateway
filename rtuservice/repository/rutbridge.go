package repository

type AcPacketRepository struct {
	SlaveId   int
	Bms       int
	Value1000 []byte
	Timer     string
}

type AcScanRepository struct {
	SlaveId   int
	Value1000 []byte
}

type AcIndoorRepository interface {
	AcScaner() ([]AcScanRepository, error)
	AcAction(int, int, int, int) (*AcPacketRepository, error)
	AcReader(int, int) (*AcPacketRepository, error)
}
