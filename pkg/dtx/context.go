package dtx

import (
	"context"

	"github.com/segmentio/ksuid"
)

const (
	RequestIDNoSet = "request id not set"
	UserIDNoSet    = -1
)

/*
If we want to add some values in context, we can define a strut like `RequestID` and functions like `WithRequestID`
and `GetRequestID` to implement it.
*/

type RequestID struct{}
type UserID struct{}
type VideoID struct{}

// NewWithRequestID used to create a context.Context and add a request id to it.
func NewWithRequestID() context.Context {
	ctx := context.Background()
	return WithRequestID(ctx)
}

// WithRequestID used to get a context.Context with a request id.
func WithRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, RequestID{}, ksuid.New().String()[0:20])
}

// GetRequestID used to get request id from a context.Context.
func GetRequestID(ctx context.Context) string {
	requestId := ctx.Value(RequestID{})
	if requestId == nil {
		return RequestIDNoSet
	}
	return requestId.(string)
}

func WithUserID(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, UserID{}, userId)
}

func GetUserID(ctx context.Context) int64 {
	userId := ctx.Value(UserID{})
	if userId == nil {
		return UserIDNoSet
	}
	return userId.(int64)
}
