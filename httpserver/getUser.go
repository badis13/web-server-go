package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	if handlerError(ctx.Writer, ctx.Request) {
		return
	}
	var user ResponseUser
	ctx.JSON(http.StatusOK, gin.H{"users": user})

}
