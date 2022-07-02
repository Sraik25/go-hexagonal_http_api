package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}
