package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
)

type AuthType string

//const available value for enum
const (
	LOGIN   AuthType = "LOGIN"
	PLAIN   AuthType = "PLAIN"
	CRAMMD5 AuthType = "CRAM-MD5"
)

//Value validate enum when set to database
func (t AuthType) Value() (driver.Value, error) {
	switch t {
	case LOGIN, PLAIN, CRAMMD5: //valid case
		return string(t), nil
	}
	return nil, errors.New("invalid auth type value") //else is invalid
}

//Scan validate enum on read from data base
func (t *AuthType) Scan(value interface{}) error {

	var at AuthType

	var st string
	switch v := value.(type) {
	case string:
		st = v
	case []uint8:

		st = string([]byte(v))
	default:
		return errors.New("invalid data for auth type: " + reflect.TypeOf(value).Name())
	}

	at = AuthType(st) //convert type from string to AuthType

	switch at {
	case LOGIN, CRAMMD5, PLAIN: //valid case
		*t = at
		return nil
	}
	return fmt.Errorf("invalid auth type value :%s", st) //else is invalid
}

type StmpCredential struct {
	BaseModel
	Name     string `gorm:"column:name; type:varchar(127); default:default; not null;uniqueIndex" json:"name"`
	Address  string `gorm:"column:address; type:varchar(127); not null" json:"address" validate:"required"`
	Username string `gorm:"column:username; type:varchar(127); not null" json:"username" validate:"required"`
	Password string `gorm:"column:password; type:varchar(127); not null" json:"password" validate:"required"`
	// AuthType AuthType `gorm:"column:auth_type; type:ENUM('LOGIN', 'PLAIN', 'CRAM-MD5'); not null;default:PLAIN" json:"auth_type"`
	AuthType AuthType `gorm:"column:auth_type; type:varchar(15);check:auth_type IN ('LOGIN', 'PLAIN', 'CRAM-MD5'); not null;default:PLAIN" json:"auth_type"`
}

func (StmpCredential) TableName() string {
	return "stmp_credential"
}

func (sc StmpCredential) Hash() string {
	return SMTPHash(sc.Address, sc.Username, sc.Password, string(sc.AuthType))
}
