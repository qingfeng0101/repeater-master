package models

import (
	"crypto/md5"
	"fmt"
)

func WeComHash(corpID string, agentID uint, corpSecret string) string {

	data := []byte(fmt.Sprintf("%s%d%s", corpID, agentID, corpSecret))

	return fmt.Sprintf("%x", md5.Sum(data))
}

func LarkHash(corpID string,  corpSecret string) string {

	data := []byte(fmt.Sprintf("%s%s", corpID,  corpSecret))

	return fmt.Sprintf("%x", md5.Sum(data))
}


func SMTPHash(address, username, password, authtype string) string {
	data := []byte(address + username + password + authtype)

	return fmt.Sprintf("%x", md5.Sum(data))
}
