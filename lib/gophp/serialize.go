package gophp

import "github.com/renjingneng/a_simple_go_project/lib/gophp/serialize"

// Serialize  Generates a storable representation of a value
func Serialize(value interface{}) ([]byte, error) {
	return serialize.Marshal(value)
}

// Unserialize Creates a PHP value from a stored representation
func Unserialize(data []byte) (interface{}, error) {
	return serialize.UnMarshal(data)
}
