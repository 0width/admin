package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var startTime = time.Now()
var strStartTime = strconv.FormatInt(startTime.Unix(), 10)

func NotModified() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "max-age=86400")
		c.Header("Last-Modified", strStartTime)

		if match := c.GetHeader("Pragma"); match == "no-cache" {
			return
		}

		if match := c.GetHeader("If-Modified-Since"); match == strStartTime {
			c.Status(http.StatusNotModified)
			c.Abort()
			return
		}
	}
}
