package rpc

import (
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

// InitRPC init rpc client
func InitRPC() {
	UserConfig := dtviper.ConfigInit("DOUTOK_USER", "user")
	initUserRpc(&UserConfig)

	FeedConfig := dtviper.ConfigInit("DOUTOK_FEED", "feed")
	initFeedRpc(&FeedConfig)

	PublishConfig := dtviper.ConfigInit("DOUTOK_PUBLISH", "publish")
	initPublishRpc(&PublishConfig)

	FavoriteConfig := dtviper.ConfigInit("DOUTOK_FAVORITE", "favorite")
	initFavoriteRpc(&FavoriteConfig)

	CommentConfig := dtviper.ConfigInit("DOUTOK_COMMENT", "comment")
	initCommentRpc(&CommentConfig)

	RelationConfig := dtviper.ConfigInit("DOUTOK_RELATION", "relation")
	initRelationRpc(&RelationConfig)

	MessageConfig := dtviper.ConfigInit("DOUTOK_MESSAGE", "message")
	initMessageRpc(&MessageConfig)
}
