package handlers

import (
	"gym/ankur/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		var LoginCredientials models.Login
		ctx.ShouldBindJSON(&LoginCredientials)
		mess, result := api.Storage.Login(LoginCredientials.Email, LoginCredientials.Password)
		if result {
			owner_id := api.Storage.GetId(LoginCredientials.Email)
			token, _ := api.GenerateJwt(owner_id)
			ctx.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": mess})
	}

}
