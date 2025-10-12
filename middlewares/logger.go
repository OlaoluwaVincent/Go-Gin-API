package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"ClientIP: %s\n"+
				"Time: %s\n"+
				"Method: %s\n"+
				"Path: %s\n"+
				"Proto: %s\n"+
				"Status: %d\n"+
				"Latency: %s\n"+
				"UserAgent: %s\n"+
				"ErrorMessage: %s\n\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
