package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Logout(ctx *gin.Context) {
	sessionID, err := ctx.Cookie("session_id")
	if err == nil {
		_ = c.sessionService.Delete(sessionID)
	}

	clearSessionCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "session terminated"})
}
