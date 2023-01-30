package pack

import (
	"context"

	"github.com/TremblingV5/DouTok/favorite/dal/db"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
)

// FavoriteVideos pack favoriteVideos info.
func FavoriteVideos(ctx context.Context, vs []db.Video, uid *int64) ([]*feed.Video, error) {
	videos := make([]*db.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}

	packVideos, err := Videos(ctx, videos, uid)
	if err != nil {
		return nil, err
	}

	return packVideos, nil
}
