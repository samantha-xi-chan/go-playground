package redislock

import (
	"fmt"
	"github.com/go-redis/redis"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"log"
)

type Locker struct {
	Client *redis.Client
	Mutex  *redsync.Mutex
}

func NewLocker(url string, name string) (*Locker, error) {

	client := goredislib.NewClient(&goredislib.Options{
		Addr: url,
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutex := rs.NewMutex(name)
	log.Printf("mutex: %#v", mutex)

	ret := Locker{
		//Client: &client,
		Mutex: mutex,
	}
	return &ret, nil
}

func (locker *Locker) Lock() error {

	if err := locker.Mutex.Lock(); err != nil {
		return fmt.Errorf("locker.Mutex.Lock : %w", err)
	}

	return nil
}

func (locker *Locker) Unlock() (bool, error) {
	ok, err := locker.Mutex.Unlock()
	if err != nil {
		return ok, fmt.Errorf("locker.Mutex.Unlock(): %w", err)
	}

	return ok, nil
}
