package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
Function: Welcome

For testing purpose and see if server can be reached or not
*/
func Welcome(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, Welcome to ITX WABizz Web Service!")
}
