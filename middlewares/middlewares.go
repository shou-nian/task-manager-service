package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() func(*gin.Context) {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.DateTime),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	})
}

func AuthRequired() func(*gin.Context) {
	return func(context *gin.Context) {
		headers := context.GetHeader("Authorization")

		if headers == "" {
			context.JSON(
				http.StatusUnauthorized,
				map[string]string{"msg": http.StatusText(http.StatusUnauthorized)},
			)

			// 中断请求
			context.Abort()
			return
		}

		// 继续请求
		context.Next()
	}
}
