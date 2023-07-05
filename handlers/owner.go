package handlers

import (
	"encoding/json"
	"gym/ankur/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (api *ApiHandler) GetOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		ListOfOwners := api.Storage.GetOwner()
		json.NewEncoder(ctx.Writer).Encode(ListOfOwners)
	}
}

func (api *ApiHandler) CreateOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		var NewOwner models.CreateOwner
		err := ctx.ShouldBindJSON(&NewOwner)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "credientials missing"})
			return
		}
		_ = api.Storage.CreateOwner(NewOwner.Name, NewOwner.Email, NewOwner.Password)
		owner_id := api.Storage.GetId(NewOwner.Email)
		token, _ := api.GenerateJwt(owner_id)
		ctx.JSON(http.StatusOK, gin.H{"token": token})

	}
}

func (api *ApiHandler) DeleteOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application.json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.DeleteOwnerFromCustomer(id)
		api.Storage.DeleteGymsOfOwners(id)
		api.Storage.DelteOwnerFromUsers(id)

	}
}
func (api *ApiHandler) UpdateOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func (api *ApiHandler) GenerateJwt(owner_id int) (string, error) {
	Key := []byte("MyGyM")
	expirationTime := time.Now().Add(time.Minute * 60).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["owner_id"] = owner_id
	claims["exp"] = expirationTime
	tokenString, _ := token.SignedString(Key)
	return tokenString, nil
}
