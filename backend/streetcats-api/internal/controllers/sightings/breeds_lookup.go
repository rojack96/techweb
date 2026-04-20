package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BreedsLookup(ctx *gin.Context) {
	animalID := ctx.Param("animalID")

	animalIDUint, err := strconv.ParseUint(animalID, 10, 64)
	if err != nil {
		c.jinres.BadRequest().Message("Invalid animal ID parameter").Done(ctx)
		return
	}

	breeds, err := c.sightingService.BreedsLookup(animalIDUint)
	if err != nil {
		c.jinres.InternalServerError().Message("Failed to lookup breed").Done(ctx)
		return
	}

	if breeds == nil {
		c.jinres.NotFound().Message("Breed not found for the given animal ID").Done(ctx)
		return
	}

	c.jinres.OK().Response(breeds).Done(ctx)
}
