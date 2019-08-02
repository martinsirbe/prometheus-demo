package postgres

import "fmt"

// InstrumentedClientOther instrumented postgres client
type InstrumentedClientOther struct{}

// NewInstrumentedClientOther initialises a new fake postgres client
func NewInstrumentedClientOther() *InstrumentedClientOther {
	return &InstrumentedClientOther{}
}

// Insert fake insert
func (c *InstrumentedClientOther) Insert(o string) error {
	fmt.Printf("posted metrics, inserted - %s\n", o)
	return nil
}

// Delete fake delete
func (c *InstrumentedClientOther) Delete(o string) error {
	fmt.Printf("posted metrics, deleted - %s\n", o)
	return nil
}
