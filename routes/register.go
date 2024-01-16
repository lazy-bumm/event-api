package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
	userId:=context.GetInt64("userId")
    
	eventId,err:=strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse event id"})
	    return
	}

	event,err:=models.GetEventById(eventId)

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"mssg":"could not fetch event"})
	    return
	}

    err =event.Register(userId)
     
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"mssg":"could not register for event"})
	    return
	}

	context.JSON(http.StatusCreated,gin.H{"mssg":"registered!"})

}


func cancelRegistration(context *gin.Context){

	userId:=context.GetInt64("userId")
    
	eventId,err:=strconv.ParseInt(context.Param("id"),10,64)
     
	var event models.Event
	event.ID=eventId

	event.CancelRegistration(userId)

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"mssg":"could not cancel registration for event"})
	    return
	}


	context.JSON(http.StatusOK,gin.H{"mssg":"cancelled"})



}