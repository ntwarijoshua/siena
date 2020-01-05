package models

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
	"time"
)

type User struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Confirmed bool      `db:"confirmed"`
	ProfileId int       `db:"profile_id"`
	RoleId    int       `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool
}

type Profile struct {
	Id           int            `db:"id"`
	Names        string         `db:"names"`
	TagLine      string         `db:"tag_line"`
	DOB          time.Time      `db:"date_of_birth"`
	ProfilePhoto sql.NullString `db:"profile_photo"`
}

var err error

func (dataStore *DataStore) SaveUser(profile Profile, user User, role Role) (User, error) {
	var (
		profileId int

		rows *sqlx.Rows

		createUserStatement = `
				INSERT INTO users (email,password,profile_id,role_id,created_at) 
				VALUES (:email,:password, :profile_id, :role_id, :created_at) RETURNING *`

		createProfileStatement = `INSERT INTO profiles (names, tag_line, date_of_birth)
					VALUES (:names, :tag_line, :date_of_birth) RETURNING id`
	)

	tx := dataStore.DB.MustBegin()
	if rows, err = tx.NamedQuery(createProfileStatement, profile); err != nil {
		logger.Errorf("Error occurred while creating a profile:", err)
	}

	for rows.Next() {
		if err = rows.Scan(&profileId); err != nil {
			logger.Errorf("Error occurred while creating profile:", err)
		}
	}
	user.ProfileId = profileId
	user.RoleId = role.Id
	user.CreatedAt = time.Now()
	if rows, err = tx.NamedQuery(createUserStatement, user); err != nil {
		logger.Errorf("Error occurred while creating a user:", err)
		return user, err
	}

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			logger.Errorf("Error occurred while creating a user:", err)
		}
	}
	err = tx.Commit()
	return user, err
}

func (dataStore *DataStore) GetUserById(userId int) (User, error) {
	var user User
	err = dataStore.DB.Get(&user, "SELECT * FROM users WHERE id = $1", userId)
	return user, err
}

func (dataStore *DataStore) GetUserByEmail(email string) (User, error) {
	var user User
	err = dataStore.DB.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}
