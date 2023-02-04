package db

import (
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/ttviper"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config = ttviper.ConfigInit("TIKTOK_DB", "dbConfig")
)

func Init() {
	InitDB()
}

// Init init DB
func InitDB() {
	var err error

	logger := dlog.NewZapGorm2(dlog.InitLog(3))
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	viper := Config.Viper
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.charset"),
		viper.GetBool("mysql.parseTime"),
		viper.GetString("mysql.loc"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			// PrepareStmt: true,
			// SkipDefaultTransaction: true,
			Logger: logger,
		},
	)
	if err != nil {
		logger.ZapLogger.Fatal(err.Error())
	}

	// gorm open telemetry records database queries and reports DBStats metrics.
	if err = DB.Use(otelgorm.NewPlugin()); err != nil {
		logger.ZapLogger.Fatal(err.Error())
	}

	// AutoMigrate会创建表，缺失的外键，约束，列和索引。如果大小，精度，是否为空，可以更改，则AutoMigrate会改变列的类型。出于保护您数据的目的，它不会删除未使用的列
	// 刷新数据库的表格，使其保持最新。即如果我在旧表的基础上增加一个字段age，那么调用autoMigrate后，旧表会自动多出一列age，值为空
	//if err := DB.AutoMigrate(&User{}, &Video{}, &Comment{}, &Relation{}); err != nil {
	//	logger.ZapLogger.Fatal(err.Error())
	//}

	sqlDB, err := DB.DB()
	if err != nil {
		logger.ZapLogger.Fatal(err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		logger.ZapLogger.Fatal(err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
