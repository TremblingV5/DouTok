// Code generated by hertz generator.

package main

import (
	"fmt"
	_ "github.com/TremblingV5/DouTok/applications/api/docs"
	"github.com/TremblingV5/DouTok/applications/api/initialize"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

//	@title			DouTokApi
//	@version		1.0
//	@description	DouTok 项目后端

//	@contact.name	DouTok
//	@contact.url	https://github.com/TremblingV5/DouTok

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8088
//	@BasePath	/
//	@schemes	http
//
// 初始化 Hertz API 及 Router
func main() {
	logger := dlog.InitHertzLog(3)
	defer logger.Sync()

	hlog.SetLogger(logger)

	initialize.Init()

	h, shutdown := initialize.InitHertz()
	defer shutdown()

	pprof.Register(h)

	register(h)

	url := swagger.URL("http://localhost:8088/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Println("http://localhost:8088/swagger/index.html")

	h.Spin()
}