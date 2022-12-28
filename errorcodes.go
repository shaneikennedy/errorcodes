package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaneikennedy/errorcodes/httpstatuscodes"
)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	for _, code := range httpstatuscodes.StatusCodes {
		r.GET("/"+strconv.Itoa(code), statusWrapper(code))
		r.POST("/"+strconv.Itoa(code), statusWrapper(code))
		r.DELETE("/"+strconv.Itoa(code), statusWrapper(code))
	}
	return r

}

func statusWrapper(code int) func(c *gin.Context) {
	return func(c *gin.Context) { c.Status(code) }
}
