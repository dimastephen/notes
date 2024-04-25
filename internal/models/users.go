package models

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type Usermodel struct {
	DB *sql.DB
}

func (m *Usermodel) Insert(name, email, password string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
		return err
	}
	stmt := "INSERT INTO users (name,email,hashed_password,created) VALUES (?,?,?,UTC_TIMESTAMP())"
	_, err = m.DB.Exec(stmt, name, email, string(hashed_password))
	if err != nil {
		var mySqlError *mysql.MySQLError
		if errors.As(err, &mySqlError) {
			if mySqlError.Number == 1062 && strings.Contains(mySqlError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (m *Usermodel) Authentificate(email, password string) (int, error) {
	return 0, nil
}

func (m *Usermodel) Exists(id int) (bool, error) {
	return false, nil
}
