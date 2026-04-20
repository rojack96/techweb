package controllers

import (
	"strconv"
	"streetcats-api/internal/dto"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BreedsLookup(ctx *gin.Context) {
	animalID := ctx.Param("animalID")

	animalIDUint, err := strconv.ParseUint(animalID, 10, 64)
	if err != nil {
		c.jinres.BadRequest().Message("Invalid animal ID parameter").Done(ctx)
		return
	}

	breed, err := c.sightingService.BreedsLookup(animalIDUint)
	if err != nil {
		c.jinres.InternalServerError().Message("Failed to lookup breed").Done(ctx)
		return
	}

	if breed == nil {
		c.jinres.NotFound().Message("Breed not found for the given animal ID").Done(ctx)
		return
	}

	response := &dto.BreedDTO{
		ID:   breed.ID,
		Name: breed.Name,
	}

	c.jinres.OK().Response(response).Done(ctx)
}
