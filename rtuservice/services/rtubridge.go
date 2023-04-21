package services

type AcIndoorInfo struct {
	SlaveId   int    `json:"slaveid"`
	Bms       int    `json:"bms"`
	Value1000 []byte `json:"value1000"`
	Timer     string `json:"timer"`
}

type AcInddorRequest struct {
	SlaveId int `json:"slaveid"`
	BmsId   int `json:"bmsid"`
	Addr    int `json:"addr"`
	Value   int `json:"value"`
}

type RtuBridgeService interface {
	GetAcValue(int, int) (*AcIndoorInfo, error)
	GetAcAction(AcInddorRequest) (*AcIndoorInfo, error)
}
