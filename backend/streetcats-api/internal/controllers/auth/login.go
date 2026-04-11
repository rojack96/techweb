package auth

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	state := generateRandomState()

	// salva lo state in cookie per verificare al ritorno
	setStateCookie(ctx, state)

	url := c.kcService.GetLoginURL(state)

	ctx.Redirect(302, url)
}
