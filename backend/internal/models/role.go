package models

import (
	_ "github.com/jmoiron/sqlx"
	"time"
)


type Role struct {
	Id int
	Name string
	Slug string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted bool
}

func (dataStore *DataStore) GetRoleBySlug(slug string) (Role, error) {
	var role  = Role{}
	err = dataStore.DB.Get(&role, "SELECT * FROM roles WHERE slug = $1", slug)
	return role, err
}
