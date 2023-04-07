package services

type AcIndoorInfo struct {
	SlaveId   int
	Bms       int
	Value1000 []byte
	Timer     string
}

type AcInddorRequest struct {
	SlaveId int
	BmsId   int
	Cmd     string
	Value   int
}

type RtuBridgeService interface {
	GetAcValue(int, int) (*AcIndoorInfo, error)
	GetAcAction(AcInddorRequest) (*AcIndoorInfo, error)
}
