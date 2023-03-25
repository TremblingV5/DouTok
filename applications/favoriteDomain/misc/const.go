package misc

var (
	FavCountTopicName = "fav_count"
	FavCountGroupName = "fav_count01"

	ViperConfigEnvPrefix   = "DOUTOK_FAVORITEDOMAIN"
	ViperConfigEnvFilename = "favoriteDomain"

	ConfigIndex_MySQLUsername = "MySQL.Username"
	ConfigIndex_MySQLPassword = "MySQL.Password"
	ConfigIndex_MySQLHost     = "MySQL.Host"
	ConfigIndex_MySQLPort     = "MySQL.Port"
	ConfigIndex_MySQLDb       = "MySQL.Database"

	ConfigIndex_RedisDest             = "Redis.Dest"
	ConfigIndex_RedisPassword         = "Redis.Password"
	ConfigIndex_RedisFavCacheDbNum    = "Redis.FavCache.Num"
	ConfigIndex_RedisFavCntCacheDbNum = "Redis.FavCntCache.Num"

	ConfigIndex_SnowFlake = "Snowflake.Node"
)
