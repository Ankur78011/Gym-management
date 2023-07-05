package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (api *ApiHandler) AuthOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tkn := ctx.GetHeader("token")
		token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected siging method")
			}
			return []byte("MyGyM"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		expireTime := int(claims["exp"].(float64))
		if expireTime < int(time.Now().Unix()) {
			ctx.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}
		if ok && token.Valid {
			value, present := claims["owner_id"]
			if !present {
				fmt.Println("no value nil")
			}
			ConvertedValue := value.(float64)
			result := api.Storage.ValidateUser(ConvertedValue)
			if !result {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			fmt.Println("Userauthorized")
			ctx.Set("owner_id", claims["owner_id"].(float64))
			ctx.Next()
		}
	}
}
