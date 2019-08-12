package service

import (
	"fmt"
	"time"

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
	errs := make(chan error, 1)

	go func() {
		for {
			if err := ll.client.Insert(fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))); err != nil {
				logrus.WithError(err).Error("stopped inserting stuff")
				errs <- err
			}
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			if err := ll.client.Delete(fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))); err != nil {
				logrus.WithError(err).Error("stopped deleting stuff")
				errs <- err
			}
			time.Sleep(time.Second * 5)
		}
	}()

	if err, open := <-errs; open {
		return err
	}

	return nil
}
