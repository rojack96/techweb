package controllers

import (
	channelhandler "sipli/notification-service/api/server/channel_handler"
	"sipli/notification-service/internal/dto"

	"github.com/gin-gonic/gin"
)

// NotificationSentEvent godoc
//
//	@Summary		Sent notification event
//	@Description	Sent a new event of notification
//	@Tags			Notification
//	@Accept			json
//	@Produce		json
//	@Param			data	body		dto.NotificationEventDTO{}										true	"Data to sent for notification"
//	@Success		200		{object}	dto.OK{response=dto.NotificationEventDTO{}}						"Data returned correctly"
//	@Success		400		{object}	dto.BadRequest{response=dto.NotificationEventDTO{}}				"Error if request is wrong"
//	@Success		401		{object}	dto.Unauthorized{}												"Unauthorized"
//	@Success		500		{object}	dto.InternalServerError{response=dto.NotificationEventDTO{}}	"Error if exists a problem to server side"
//	@Security		BasicAuth
//	@Router			/notification/event [post]
func (c *Controller) NotificationSentEvent(ctx *gin.Context, r *channelhandler.ChannelHandler) {
	var (
		emailInfo []dto.EmailInfoDTO
		request   dto.NotificationEventDTO
		err       error
	)

	if err = ctx.ShouldBindJSON(&request); err != nil {
		c.log.Error("failed to bind JSON payload")
		c.jinres.BadRequest().Done(ctx)
		return
	}

	c.notificationService.SetContext(ctx)

	if emailInfo, err = c.notificationService.SaveEvent(r, request); err != nil {
		c.log.Error("failed to save event")
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	// Pubblica job su Redis
	c.emailService.SetContext(ctx)
	c.emailService.SentEmailOnRedis(emailInfo)

	c.jinres.OK().Message("OK").Done(ctx)
}
