package handlers

import (
	"encoding/json"
	"fmt"
	"gym/ankur/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) GetGyms() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		owner_id := ctx.GetFloat64("owner_id")
		fmt.Println(owner_id)
		ListOfGyms := api.Storage.GetGyms(owner_id)
		json.NewEncoder(ctx.Writer).Encode(ListOfGyms)
	}
}

func (api *ApiHandler) CreateGyms() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		owner_id := int(ctx.GetFloat64("owner_id"))
		var NewGym models.Gym
		ctx.ShouldBindJSON(&NewGym)
		fmt.Println(owner_id, NewGym.Owner_id)
		if owner_id == NewGym.Owner_id {
			mess := api.Storage.CreateGym(NewGym.Name, NewGym.Address, int(owner_id))
			ctx.JSON(http.StatusOK, gin.H{"message": mess})
			return
		}
		ctx.AbortWithStatus(http.StatusUnauthorized)

	}
}

func (api *ApiHandler) DeleteGyms() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		owner_id := int(ctx.GetFloat64("owner_id"))
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		ctx.Header("Content-Type", "application.json")
		api.Storage.DeleteGyms(id, owner_id)
	}
}
func (api *ApiHandler) UpdateGyms() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
