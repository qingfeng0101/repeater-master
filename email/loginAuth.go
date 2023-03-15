package email

//github.com/go-gomail/gomail/blob/master/auth.go
import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
)

type loginAuth struct {
	username string
	password string
	host     string
}

//LoginAuth return a implemented by smtp.Auth of LOGIN mechanism
func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "LOGIN" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("email: unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, errors.New("email: wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}

	switch {
	case bytes.Equal(fromServer, []byte("Username:")):
		return []byte(a.username), nil
	case bytes.Equal(fromServer, []byte("Password:")):
		return []byte(a.password), nil
	default:
		return nil, fmt.Errorf("email: unexpected server challenge: %s", fromServer)
	}
}
