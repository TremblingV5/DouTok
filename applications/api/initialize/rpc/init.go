package rpc

import (
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

// InitRPC init rpc client
func InitRPC() {
	UserConfig := dtviper.ConfigInit("DOUTOK_USER", "user", nil)
	initUserRpc(UserConfig)

	FeedConfig := dtviper.ConfigInit("DOUTOK_FEED", "feed", nil)
	initFeedRpc(FeedConfig)

	PublishConfig := dtviper.ConfigInit("DOUTOK_PUBLISH", "publish", nil)
	initPublishRpc(PublishConfig)

	FavoriteConfig := dtviper.ConfigInit("DOUTOK_FAVORITE", "favorite", nil)
	initFavoriteRpc(FavoriteConfig)

	CommentConfig := dtviper.ConfigInit("DOUTOK_COMMENT", "comment", nil)
	initCommentRpc(CommentConfig)

	RelationConfig := dtviper.ConfigInit("DOUTOK_RELATION", "relation", nil)
	initRelationRpc(RelationConfig)

	MessageConfig := dtviper.ConfigInit("DOUTOK_MESSAGE", "message", nil)
	initMessageRpc(MessageConfig)
}
