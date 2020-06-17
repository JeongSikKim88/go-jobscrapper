package mydict

import (
	"errors"
)

var errNotFound = errors.New("Not Found")

// Dictionary type
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}