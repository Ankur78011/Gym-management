package handlers

import (
	"encoding/json"
	"gym/ankur/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) GetReports() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		cid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		Report := api.Storage.GetReports(cid)
		json.NewEncoder(ctx.Writer).Encode(Report)
	}
}
func (api *ApiHandler) CreateReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		cid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		response := api.Storage.CheckCustomer(cid)
		if response {
			var newRecord models.NewReport
			err = ctx.ShouldBindJSON(&newRecord)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "information missing"})
			}
			api.Storage.InsertReports(cid, newRecord.Weight)
			return
		}
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}
