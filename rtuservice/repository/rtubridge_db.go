package repository

type IndoorInfoDB struct {
	ID       string      `db:"id"`
	Uid      int         `db:"uid"`
	Bms      int         `db:"bms"`
	DeviceSn string      `db:"devicesn"`
	Acvalue  indoorValue `db:"acvalue"`
}

type indoorValue struct {
	Value1000 []byte `db:"value1000"`
	Value2000 []byte `db:"value2000"`
	Value3000 []byte `db:"value3000"`
	Value4000 []byte `db:"value4000"`
}

type AcRepositoryDB interface {
	AcNewer(IndoorInfoDB) (*IndoorInfoDB, error)
	AcUpdater(int, IndoorInfoDB) (bool, error)
	AcDeleter(int) (bool, error)
	AcReader(int) (*IndoorInfoDB, error)
}
