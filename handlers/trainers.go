package handlers

import (
	"encoding/json"
	"gym/ankur/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) GetTrainer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application.json")
		ListOfTrainers := api.Storage.GetTrainers()
		json.NewEncoder(ctx.Writer).Encode(ListOfTrainers)
	}
}

func (api *ApiHandler) CreateTrainer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Constent-type", "application.json")
		var NewTrainer models.Trainer
		ctx.ShouldBindJSON(&NewTrainer)
		api.Storage.CreateTrainers(NewTrainer.Name)
	}
}

func (api *ApiHandler) DeleteTrainer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Constent-type", "application.json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.DeleteTrainer(id)

	}
}
func (api *ApiHandler) UpdateTrainer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
