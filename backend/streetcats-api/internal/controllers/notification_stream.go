package controllers

import (
	"fmt"
	"io"
	channelhandler "sipli/notification-service/api/server/channel_handler"
	"sipli/notification-service/internal/dto"
	"time"

	"github.com/gin-gonic/gin"
)

// NotificationStream godoc
//
//	@Summary		Stream notification
//	@Description	Stream of new notification
//	@Tags			Notification
//	@Accept			mpfd
//	@Produce		json
//	@Success		200	{object}	dto.OK{}					"Data returned correctly"
//	@Success		400	{object}	dto.BadRequest{}			"Error if request is wrong"
//	@Success		401	{object}	dto.Unauthorized{}			"Unauthorized"
//	@Success		500	{object}	dto.InternalServerError{}	"Error if exists a problem to server side"
//	@Security		Bearer
//	@Router			/notification/stream [get]
func (c *Controller) NotificationStream(ctx *gin.Context, r *channelhandler.ChannelHandler) {
	var (
		claims  any
		history []dto.NotificationDTO

		userId uint64

		msg        dto.NotificationDTO
		ch         chan dto.NotificationDTO
		unregister func()

		ok  bool
		err error
	)

	if claims, ok = ctx.Get("claims"); !ok {
		c.log.Error("failed to get claims from context")
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	if userId, err = c.notificationService.GetUserIdByPreferredUsername(claims); err != nil {
		c.jinres.InternalServerError().Message(err.Error()).Done(ctx)
		return
	}

	if ch, unregister, err = r.Register(userId); err != nil {
		c.jinres.TooManyRequests().Done(ctx)
		return
	}
	defer unregister()

	// Sent history
	if history, err = c.notificationService.GetNotifications(userId); err != nil {
		c.jinres.InternalServerError().Message(err.Error()).Done(ctx)
		return
	}

	/*for _, historyMsg := range history {
		ctx.SSEvent("notification-history", historyMsg)
	}*/
	ctx.SSEvent("notification-history", history)
	// forza flush immediato
	ctx.Writer.Flush()

	heartbeat := time.NewTicker(15 * time.Second)
	defer heartbeat.Stop()

	// Real time SSE
	ctx.Stream(func(w io.Writer) bool {
		select {
		case msg, ok = <-ch:
			if !ok {
				return false
			}

			ctx.SSEvent("notification-event", msg)
			ctx.Writer.Flush()
			return true
		case <-heartbeat.C:
			// comment SSE = keep-alive
			fmt.Fprint(w, ": ping\n\n")
			ctx.Writer.Flush()
			return true

		case <-ctx.Request.Context().Done():
			return false
		}
	})

}

/*

-------------------- Functions and utility types  --------------------
Here are the structures, functions useful for the operation of Raw Data Logs

*/
