package recovery

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[Middleware] %s panic recovered:\n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), err)

				ctx.Abort()
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		ctx.Next()
	}
}
