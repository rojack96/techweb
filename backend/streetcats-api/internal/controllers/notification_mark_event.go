package controllers

import (
	"sipli/notification-service/internal/dto"

	"github.com/gin-gonic/gin"
)

// NotificationMarkEvent godoc
//
//	@Summary		Mark read notification
//	@Description	Mark notification as read or not
//	@Tags			Notification
//	@Accept			mpfd
//	@Produce		json
//	@Param			data	body		dto.NotificationMarkerDTO{}										true	"Marked notification read or not"
//	@Success		200		{object}	dto.OK{response=dto.NotificationMarkerDTO{}}					"Data returned correctly"
//	@Success		400		{object}	dto.BadRequest{response=dto.NotificationMarkerDTO{}}			"Error if request is wrong"
//	@Success		401		{object}	dto.Unauthorized{response=dto.NotificationMarkerDTO{}}			"Unauthorized"
//	@Success		500		{object}	dto.InternalServerError{response=dto.NotificationMarkerDTO{}}	"Error if exists a problem to server side"
//	@Security		Bearer
//	@Router			/notification/mark-event [patch]
func (c *Controller) NotificationMarkEvent(ctx *gin.Context) {
	var (
		claims any

		request dto.NotificationMarkerDTO

		ok  bool
		err error
	)

	if err = ctx.ShouldBindJSON(&request); err != nil {
		c.log.Error("failed to bind JSON payload")
		c.jinres.BadRequest().Done(ctx)
		return
	}

	if claims, ok = ctx.Get("claims"); !ok {
		c.log.Error("failed to get claims from context")
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	if err = c.notificationService.UpdateEventByUserId(claims, request.IDs, request.Marked); err != nil {
		c.log.Error("failed to update event by user")
		c.jinres.BadRequest().Done(ctx)
		return
	}

	c.jinres.OK().Message("event updated correctly").Done(ctx)
	return
}
