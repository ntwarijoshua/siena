package models

import "github.com/jmoiron/sqlx"

type DataLayer interface {

}
type DataStore struct {
	DB *sqlx.DB
}



