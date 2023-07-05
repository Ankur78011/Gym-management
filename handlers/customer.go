package handlers

import (
	"encoding/json"
	"gym/ankur/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) GetCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		owner_id := int(ctx.GetFloat64("owner_id"))
		ctx.Header("Content-Type", "application.json")
		ListOfCustomers := api.Storage.GetCustomerDetails(owner_id)
		json.NewEncoder(ctx.Writer).Encode(ListOfCustomers)

	}
}

func (api *ApiHandler) CreateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Contnent-type", "application.json")
		var NewCustomer models.CreateCustomer
		ctx.ShouldBindJSON(&NewCustomer)
		id := api.Storage.InserCustomer(NewCustomer.Name)
		api.Storage.WeightSection(NewCustomer.Current_weight, NewCustomer.Targeted_Weight, id)
		api.Storage.Personal_trainer(NewCustomer.Personal_trainer, id)
		api.Storage.GymId(NewCustomer.Gym_Name, id)
		ctx.JSON(http.StatusOK, gin.H{"message": "customer created"})
	}
}

func (api *ApiHandler) DeleteCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		owner_id := int(ctx.GetFloat64("owner_id"))
		ctx.Header("Contnent-type", "application.json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.DeleteCustomer(id, owner_id)
	}
}
func (api *ApiHandler) UpdateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
