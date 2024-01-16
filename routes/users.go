package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err:=context.ShouldBindJSON(&user)
   
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse"})
		return 
	}

	err=user.Save()
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not save user"})
		return 
	}

	context.JSON(http.StatusCreated,gin.H{"mssg":"user created successfully"})

}

func login(context *gin.Context)  {
	var user models.User
	err:=context.ShouldBindJSON(&user)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse"})
		return 
	}
    
	err=user.ValidateCredentials()
	if err!=nil{
		context.JSON(http.StatusUnauthorized,gin.H{"mssg":"Could not authenticate"})
		return
	}

	
     
	token,err:=utils.GenerateToken(user.Email,user.ID)
	
	if err!=nil{
		context.JSON(http.StatusUnauthorized,gin.H{"mssg":"Could not authenticate"})
		return
	}
   context.JSON(http.StatusOK,gin.H{"mssg":"logged in successfully","token":token})

}