package initialize

import "github.com/TremblingV5/DouTok/pkg/dtviper"

var (
	ViperConfig *dtviper.Config
)

func InitViper() {
	ViperConfig = dtviper.ConfigInit("DOUTOK_API", "api", nil)
}
