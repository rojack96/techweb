package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) AllSightings(ctx *gin.Context) {
	animalID := ctx.Param("animalID")
	if animalID == "" {
		c.log.Error("animalID parameter is required")
		c.jinres.BadRequest().Message("Animal ID parameter is required").Done(ctx)
		return
	}

	sightings, err := c.sightingService.GetAllSightings(ctx.Request.Context())
	if err != nil {
		c.log.Error("failed to get all sightings", zap.Error(err))
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	c.jinres.OK().Custom("sightings", sightings).Done(ctx)
}
