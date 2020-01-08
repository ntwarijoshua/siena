package models

import (
	"database/sql"
)

type DataLayer interface {

}
type DataStore struct {
	DB *sql.DB
}



