package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewTestHandler() *Handler {
	service.Init()
	return New(service.New())
}

func TestRelationDomainHandler_ListRelation(t *testing.T) {
	ctx := context.Background()
	handler := NewTestHandler()

	t.Run("ListFollowList", func(t *testing.T) {
		req := &relationDomain.DoutokListRelationRequest{
			ActionType: 0,
			UserId:     rand.Int63(),
		}

		resp, err := handler.ListRelation(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		// assert.Equal(t, followList, resp.UserList)
		assert.Equal(t, errno.SuccessCode, resp.StatusCode)
	})

	t.Run("ListFollowerList", func(t *testing.T) {
		req := &relationDomain.DoutokListRelationRequest{
			ActionType: 1,
			UserId:     rand.Int63(),
		}

		resp, err := handler.ListRelation(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		//assert.Equal(t, followerList, resp.UserList)
		assert.Equal(t, errno.SuccessCode, resp.StatusCode)
	})

	t.Run("ListFriendList", func(t *testing.T) {
		req := &relationDomain.DoutokListRelationRequest{
			ActionType: 2,
			UserId:     rand.Int63(),
		}

		resp, err := handler.ListRelation(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		//assert.Equal(t, friendList, resp.UserList)
		assert.Equal(t, errno.SuccessCode, resp.StatusCode)
	})
}
