package prior

import (
	"errors"
	"log"
	"testing"
	"time"
)

type BootConfig struct{}

func (b BootConfig) Initialize() error {
	time.Sleep(time.Second)
	log.Println("Config Initialized")
	return nil
}

type BootMySQL struct{}

func (b BootMySQL) Initialize() error {
	time.Sleep(time.Second)
	log.Println("MySQL Initialized")
	return nil
}

type BootSphinx struct{}

func (b BootSphinx) Initialize() error {
	time.Sleep(2 * time.Second)
	return errors.New("sphinx timeout")
}

type BootRedis struct{}

func (b BootRedis) Initialize() error {
	time.Sleep(time.Second)
	log.Println("Redis Initialized")
	return nil
}

func TestPrior_Boot(t *testing.T) {
	var bootConfig BootConfig
	var bootMysql BootMySQL
	var bootRedis BootRedis

	Register(bootConfig).
		Register(bootMysql, bootRedis)

	err := Boot()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(err)
}

func TestPrior_BootErr(t *testing.T) {
	var bootConfig BootConfig
	var bootSphinx BootSphinx
	var bootRedis BootRedis

	Register(bootConfig).
		Register(bootSphinx, bootRedis)

	err := Boot()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(err)
}
