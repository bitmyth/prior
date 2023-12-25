package prior

import (
	"log"
	"sync"
)

var (
	BootRoot = new(Prior)
)

type Initializer interface {
	Initialize() error
}

type Prior struct {
	Initializers []Initializer
	Then         *Prior
	sync.RWMutex
}

func Register(initializer ...Initializer) *Prior {
	return BootRoot.Register(initializer...)
}

func (b *Prior) Register(initializer ...Initializer) *Prior {
	b.Lock()
	defer b.Unlock()

	b.Initializers = append(b.Initializers, initializer...)
	b.Then = new(Prior)

	return b.Then
}

func Boot() error {
	return BootRoot.boot()
}

func (b *Prior) boot() error {
	if len(b.Initializers) == 0 {
		return nil
	}

	resultCh := make(chan error, len(b.Initializers))

	for _, initializer := range b.Initializers {
		initializer := initializer
		go func() {
			resultCh <- initializer.Initialize()
		}()
	}

	count := 0

	for err := range resultCh {
		count++
		log.Println("result: ", err)
		if err != nil {
			log.Println("error:", err)
			return err
		}

		if count == len(b.Initializers) {
			break
		}
	}

	if b.Then == nil {
		return nil
	}

	return b.Then.boot()
}
