package ex1

import (
	"encoding/json"
	"errors"
)

type MapStore map[ID]*json.RawMessage

func (m MapStore) Insert(r Record, id ID) {
	// делаем приведение типов,
	// освобождая от этого вызывающего
	m[id] = r.(*json.RawMessage)
}
func (m MapStore) Get(id ID) (*json.RawMessage, error) {
	r, ok := m[id]
	// проверяем, есть ли запись в хранилище
	if !ok {
		return r, errors.New("not found")
	}

	return r, nil

}

// конструктор
func NewMapStore() MapStore {
	s := make(MapStore)
	return s
}
