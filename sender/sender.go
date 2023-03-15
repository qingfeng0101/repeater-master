package sender

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/kexirong/repeater/models"
)

type Sender interface {
	Type() string
	Name() string
	Hash() string



	Send(to []models.Contact, subject, contentType string, content *json.RawMessage) error
}

var senderRepos *Senders

func SenderRepos() *Senders {
	return senderRepos
}

type Senders struct {
	mtx   sync.Mutex
	repos map[string]Sender
}

func (s *Senders) Register(name string, sender Sender) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.repos[name]; ok {
		return errors.New(name + ": exist")
	}
	s.repos[name] = sender
	return nil
}

func (s *Senders) List() map[string]Sender {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repos
}

func (s *Senders) Len() int {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return len(s.repos)
}

func (s *Senders) Get(name string) Sender {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repos[name]
}

func (s *Senders) Delete(name string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.repos, name)
}

func NewSenders() *Senders {
	return &Senders{
		repos: make(map[string]Sender),
	}
}

func init() {
	senderRepos = NewSenders()
}
