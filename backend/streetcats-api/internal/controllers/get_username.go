package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) GetUsernameByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		c.log.Error("email parameter is required")
		c.jinres.BadRequest().Message("Email parameter is required").Done(ctx)
		return
	}

	username, err := c.usersService.GetUsernameByEmail(ctx.Request.Context(), email)
	if err != nil {
		c.log.Error("failed to get username by email", zap.Error(err))
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	c.jinres.OK().Custom("username", username).Done(ctx)
}
