package controllers

import (
	"streetcats-api/internal/dto"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateSighting(ctx *gin.Context) {
	var sightingDTO dto.CreateSightingDTO

	if err := ctx.ShouldBindJSON(&sightingDTO); err != nil {
		c.jinres.BadRequest().Message("Invalid request body").Done(ctx)
		return
	}

	sightingID, err := c.sightingService.CreateSighting(sightingDTO)
	if err != nil {
		c.jinres.InternalServerError().Message("Failed to create sighting").Done(ctx)
		return
	}

	sightingDTO.ID = sightingID

	c.jinres.Created().Response(sightingDTO).Done(ctx)
}
