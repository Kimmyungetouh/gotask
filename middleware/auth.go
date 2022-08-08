package middleware

import (
	"TaskManager/helpers"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware( /*context *gin.Context*/ ) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helpers.TokenValid(context)
		helpers.HandleError(context, "Token is not valid", err)
	}
	/*err := helpers.TokenValid(context)
	helpers.HandleError(context, "Token is not valid", err)
	context.Next()*/
}
