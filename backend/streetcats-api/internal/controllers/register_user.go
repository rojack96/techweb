package controllers

import (
	"streetcats-api/internal/dto"

	"github.com/gin-gonic/gin"
)

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var (
		request dto.AccountDTO
		err     error
	)

	if err = ctx.ShouldBindJSON(&request); err != nil {
		c.log.Error("failed to bind JSON payload")
		c.jinres.BadRequest().Done(ctx)
		return
	}

	err = c.usersService.CreateUser(ctx.Request.Context(), request)
	if err != nil {
		c.log.Error("failed to create user")
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	c.jinres.Created().Message("User registered").Done(ctx)
}
