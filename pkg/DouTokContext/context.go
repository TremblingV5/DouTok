package DouTokContext

import (
	"context"
	"github.com/segmentio/ksuid"
)

const (
	RequestIDNoSet = "request id not set"
)

/*
If we want to add some values in context, we can define a strut like `RequestID` and functions like `WithRequestID`
and `GetRequestID` to implement it.
*/

type RequestID struct{}

func New() context.Context {
	ctx := context.Background()
	return WithRequestID(ctx)
}

func WithRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, RequestID{}, ksuid.New().String()[0:20])
}

func GetRequestID(ctx context.Context) string {
	requestId := ctx.Value(RequestID{})
	if requestId == nil {
		return RequestIDNoSet
	}
	return requestId.(string)
}
