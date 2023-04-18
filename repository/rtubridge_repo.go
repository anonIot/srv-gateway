package repository

import (
	"github.com/jmoiron/sqlx"
)

type acIndoorDB struct {
	db *sqlx.DB
}

func NewIndoorDB(db *sqlx.DB) AcRepositoryDB {
	return acIndoorDB{db: db}
}

func (r acIndoorDB) AcNewer(acInfo IndoorInfoDB) (*IndoorInfoDB, error) {
	return nil, nil
}
func (r acIndoorDB) AcUpdater(id int, acInfo IndoorInfoDB) (bool, error) {
	return false, nil
}
func (r acIndoorDB) AcDeleter(id int) (bool, error) {
	return false, nil
}
func (r acIndoorDB) AcReader(id int) (*IndoorInfoDB, error) {
	return nil, nil
}
