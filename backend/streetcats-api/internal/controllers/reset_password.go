package controllers

import (
	"streetcats-api/internal/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) ResetPassword(ctx *gin.Context) {
	var (
		request dto.ResetPasswordDTO
		err     error
	)

	if err = ctx.ShouldBindJSON(&request); err != nil {
		c.log.Error("failed to bind JSON payload")
		c.jinres.BadRequest().Done(ctx)
		return
	}

	err = c.usersService.ResetPassword(ctx.Request.Context(), request.Identifier, request.NewPassword)
	if err != nil {
		c.log.Error("failed to reset password", zap.Error(err))
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	c.jinres.OK().Message("Password reset successfully").Done(ctx)
}
