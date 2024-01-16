package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"

	"github.com/gin-gonic/gin"
)




func getEvents(context *gin.Context) {
	events,err:=models.GetAllEvents()

	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"mssg":"could not fetch events"})
	    return
	}
    context.JSON(http.StatusOK,events)
}

func getEvent(context *gin.Context){
	eventId,err:=strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse event id"})
	    return
	}

	event,err:=models.GetEventById(eventId)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not get event with id"})
	    return
	}

	context.JSON(http.StatusOK,event)


}

func createEvent(context *gin.Context){



    var event models.Event

	//similar to scan 
	err:=context.ShouldBindJSON(&event)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse"})
		return 
	}
   userId:=context.GetInt64("userId")
	
	event.UserID=userId


	event.Save()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"mssg":"could not create event"})
	    return
	}
	context.JSON(http.StatusCreated,gin.H{"mssg":"event created","event":event})


}

func updateEvent(context *gin.Context){

    userId:=context.GetInt64("userId")
	eventId,err:=strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse event id"})
	    return
	}

	event, err:=models.GetEventById(eventId)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not get event with id"})
	    return
	}

	if event.UserID!=userId{
		context.JSON(http.StatusUnauthorized,gin.H{"mssg":"not authorized to update event"})
	    return
	}

	var updatedEvent models.Event

	err=context.ShouldBindJSON(&updatedEvent)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse req data updated"})
	    return
	}
   
	updatedEvent.ID=eventId
    err=updatedEvent.Update()

	if err !=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not update event"})
	    return
	}
	context.JSON(http.StatusOK,gin.H{"mssg":"Successfully Event updated"})
}

func deleteEvent(context *gin.Context){


	userId:=context.GetInt64("userId")
	eventId,err:=strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not parse event id"})
	    return
	}

	event, err:=models.GetEventById(eventId)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not get event with id"})
	    return
	}

	if event.UserID!=userId{
		context.JSON(http.StatusUnauthorized,gin.H{"mssg":"not authorized to delete event"})
	    return
	}

   
	err=event.Delete()

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"mssg":"could not delete event"})
	    return
	}

	context.JSON(http.StatusBadRequest,gin.H{"mssg":"event deleted successfully"})
	  

}