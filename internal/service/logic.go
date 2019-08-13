package service

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -package=mocks -destination=../mocks/storage.go github.com/martinsirbe/prometheus-demo/internal/service Storage

// Storage dummy storage interface used for demo
type Storage interface {
	Insert(o string) error
	Delete(o string) error
}

// LogicLayer logic implementation which uses storage client
type LogicLayer struct {
	client Storage
}

// NewLogicLayer initialises a new service logic layer
func NewLogicLayer(client Storage) *LogicLayer {
	return &LogicLayer{client}
}

// Run runs demo
func (ll *LogicLayer) Run() error {
	for {
		var (
			maxThreads = 5
			wg         sync.WaitGroup
		)

		wg.Add(maxThreads)
		for i := 0; i < maxThreads; i++ {
			go func() {
				if err := ll.client.Insert(id()); err != nil {
					logrus.WithError(err).Error("failed to insert")
				}

				if err := ll.client.Delete(id()); err != nil {
					logrus.WithError(err).Error("failed to delete")
				}

				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func id() string {
	return fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))
}
