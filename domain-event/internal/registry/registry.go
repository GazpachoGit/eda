package registry

import (
	"encoding/json"
)

type (
	Registry interface {
		Serialize(v interface{}) ([]byte, error)
		Deserialize(data []byte) (interface{}, error)
	}
)

type registry struct {
}

func NewRegistry() Registry {
	return registry{}
}

func (r registry) Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (r registry) Deserialize(data []byte) (interface{}, error) {
	var res interface{}
	err := json.Unmarshal(data, res)
	return res, err
}
