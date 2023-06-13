package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var input ResponseUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Сonnection.Exec(`INSERT INTO human (id, name, age)
    VALUES(?, ?, ?)`, input.Id, input.Name, input.Age)

	ctx.JSON(http.StatusOK, gin.H{})

}
