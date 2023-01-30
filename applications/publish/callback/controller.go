package callback

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CallbackDto struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

func CallbackAction(c *gin.Context) {
	var in CallbackDto
	if err := c.Bind(&in); err != nil {
		fmt.Println(err)
	}
	fmt.Println(in)

	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
