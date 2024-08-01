package response

import (
	"github.com/gin-gonic/gin"
	"go-learning/3-gorm/common/result"
	"net/http"
)

func Ok(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, result.Ok())
}

func OkWithData(ctx *gin.Context, data result.Any) {
	ctx.JSON(http.StatusOK, result.OkWithData(data))
}

func Error(ctx *gin.Context, code uint16, message string) {
	ctx.JSON(http.StatusOK, result.Error(code, message))
}

func Json(ctx *gin.Context, code uint16, message string, data result.Any) {
	ctx.JSON(http.StatusOK, result.Of(code, message, data))
}
