package store

import "errors"

var ErrEmptykey = errors.New("key must not be empty")

type Storer interface {
	Get(key string) (string, bool,error)
	Set(key, value string) error
	Keys() []string
	Delete(key string)
	Len() int
	Incr(key string) (int,error)
	
}
