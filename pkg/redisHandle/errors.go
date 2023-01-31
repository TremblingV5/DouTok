package redishandle

import "errors"

type RedisError struct {
	Err    error
	Supple string
}

func NewError(err error, supple string) *RedisError {
	return &RedisError{
		Err:    err,
		Supple: supple,
	}
}

func (e *RedisError) String() string {
	return e.Err.Error() + ";" + e.Supple
}

var (
	ErrNotEnoughOpNumsInList = errors.New("there are not enough op times in the list")
)
