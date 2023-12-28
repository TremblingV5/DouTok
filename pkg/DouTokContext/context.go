package DouTokContext

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
