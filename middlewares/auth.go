package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token:=context.Request.Header.Get("Authorization")

	if token ==""{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"mssg":" not authorize"})
		return 
	}

	userId,err:=utils.VerifyToken(token)

	if err!=nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"mssg":" not authorize"})
		return 
	}


     context.Set("userId",userId)
	
// next req handler in line will execute
	context.Next()


}