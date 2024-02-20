package services

import (
	"github.com/cloudwego/kitex/client"
)

type Service[T any] struct {
	Client T
}

func New[T any](name string, op func(destService string, opts ...client.Option) (T, error), options []client.Option) *Service[T] {
	client, err := op(name, options...)
	if err != nil {
		panic(err)
	}
	return &Service[T]{Client: client}
}
