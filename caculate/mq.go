package caculate

import (
	"errors"

	"golang.org/x/sync/syncmap"
)

type Mq struct {
	syncmap.Map
}

func NewMq() *Mq {
	m := new(Mq)
	return m
}
func (m *Mq) Pub(data StaInfo) {
	b := func(key, vaule interface{}) bool {
		dc := vaule.(chan StaInfo)
		select {
		case dc <- data:

		default:

		}
		return true
	}
	m.Range(b)
}
func (m *Mq) Sub(id string, data chan StaInfo) error {
	_, ok := m.Load(id)
	if ok {
		errors.New("the id stored")
	} else {
		m.Store(id, data)
	}
	return nil
}
