package controllers

import (
	EventMgr "../business"
	EventTypes "../types"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

/*
EventController struct
*/
type EventController struct {
	eventProvider EventMgr.EventProvider
}

func (ec EventController) getEventList(context *gin.Context) {
	var result, err = ec.eventProvider.GetEventList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func (ec EventController) getEvent(context *gin.Context) {
	id := context.Param("id")
	var result, err = ec.eventProvider.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func (ec EventController) insertEvent(context *gin.Context) {
	var event EventTypes.Event
	context.BindJSON(&event)
	var result = ec.eventProvider.InsertEvent(event)
	if result == false {
		context.JSON(http.StatusInternalServerError, nil)
	} else {
		context.JSON(http.StatusOK, event)
	}
}

/*
Start EventController controller
*/
func (ec EventController) Register(router *gin.Engine) {
	router.GET("/events", ec.getEventList)
	router.GET("/events/:id", ec.getEvent)
	router.POST("/events", ec.insertEvent)
}
