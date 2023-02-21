package rpc

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitPRCClient() {
	config := dtviper.ConfigInit("DOUTOK_FEED", "feed")

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.Viper.GetString("Server.Name")),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	InitCommentRpc()
	InitFavoriteRpc()
	InitUserRpc()
}
