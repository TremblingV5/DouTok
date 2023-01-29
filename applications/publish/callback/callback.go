package callback

import (
	"github.com/gin-gonic/gin"
)

func InitCallbackServer() {
	server := gin.Default()

	router := server.Group("/callback")
	router.POST("/handle", CallbackAction)

	server.Run(":8888")
}
