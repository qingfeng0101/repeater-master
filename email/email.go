package email

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/gomarkdown/markdown"

	"github.com/gomarkdown/markdown/parser"
	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
)

//SMTP smtp auth info
type SMTP struct {
	name,
	address,
	username,
	password,
	authtype string
}

//New return *SMTP
func New(name, address, username, password, authtype string) *SMTP {
	return &SMTP{
		name:     name,
		address:  address,
		username: username,
		password: password,
		authtype: authtype,
	}
}

//SendMail sent the email with args
func (s *SMTP) SendMail(from string, to []string, subject, contentType string, body []byte) error {

	tos := strings.Join(to, ";")
	addrArr := strings.Split(s.address, ":")
	if len(addrArr) != 2 {
		return fmt.Errorf("address format error")
	}

	//base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	b64 := base64.StdEncoding
	header := make(map[string]string)

	header["From"] = s.username
	suffix := strings.Split(s.username, "@")[1]
	if strings.HasSuffix(from, suffix) {
		header["From"] = from
	}
	header["To"] = tos
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"

	switch contentType {
	case "markdown":

		markdownParser := parser.NewWithExtensions(parser.CommonExtensions | parser.HardLineBreak)

		// htmlRenderer := html.NewRenderer(html.RendererOptions{
		// 	Flags: html.CommonFlags | html.TOC,
		// })
		body = markdown.ToHTML(body, markdownParser, nil)
		fallthrough
	case "html":
		header["Content-Type"] = "text/html; charset=UTF-8"
	default:
		header["Content-Type"] = "text/plain; charset=UTF-8"
	}

	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString(body)
	var auth smtp.Auth
	switch s.authtype {
	case "LOGIN":
		auth = LoginAuth(s.username, s.password, addrArr[0])
	case "CRAM-MD5":
		auth = smtp.CRAMMD5Auth(s.username, s.password)
	case "PLAIN":
	default:
		auth = smtp.PlainAuth("", s.username, s.password, addrArr[0])
	}

	return smtp.SendMail(s.address, auth, s.username, to, []byte(message))
}

// Send implement interface Sender
func (SMTP) Type() string {
	return "email"
}

func (s SMTP) Name() string {
	return fmt.Sprintf("%s-%s", s.Type(), s.name)
}

func (s SMTP) Hash() string {
	return models.SMTPHash(s.address, s.username, s.password, s.authtype)
}

func (s *SMTP) Send(contacts []models.Contact, subject, contentType string, content *json.RawMessage) error {
	var to []string
	for _, contact := range contacts {
		suffix := strings.Split(s.username, "@")[1]
		if contact.Email == "" {
			contact.Email = contact.Username + "@" + suffix
		}
		to = append(to, contact.Email)
	}
	var body string
	if err := json.Unmarshal(*content, &body); err != nil {
		body = string(*content)
	}

	return s.SendMail("", to, subject, contentType, []byte(body))
}

func StmpCredentialRegist(sc models.StmpCredential) {
	stmp := New(sc.Name, sc.Address, sc.Username, sc.Password, string(sc.AuthType))
	// name := fmt.Sprintf("%s-%s", stmp.Type(), sc.Name)
	sender.SenderRepos().Register(stmp.Name(), stmp)
}
