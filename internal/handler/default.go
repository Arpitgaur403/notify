package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultGetSlash(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Default Get request",
	})

}

func DefaultPostSlash(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Default Post request",
	})
}
