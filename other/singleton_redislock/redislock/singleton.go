package redislock

import (
	"fmt"
	"log"
	"sync"
)

var instProd *MySingleton
var onceProd sync.Once

func GetInstance() *MySingleton {
	onceProd.Do(func() {
		instProd = &MySingleton{}
	})

	return instProd
}

type MySingleton struct {
	data int
	pMQ  *Locker
}

func (s *MySingleton) Init(url, name string) error {
	mq, err := NewLocker(url, name)
	if err != nil {
		return fmt.Errorf("NewRabbitMQ : %w", err)
	}
	s.pMQ = mq
	log.Println("NewLocker ok")

	return nil
}

func (s *MySingleton) Lock() error {
	e := s.pMQ.Lock()
	if e != nil {
		return fmt.Errorf("s.pMQ.Publish : %w", e)
	}

	return nil
}

func (s *MySingleton) Unlock() (bool, error) {
	ok, e := s.pMQ.Unlock()
	return ok, e
}
