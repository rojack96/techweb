package auth

import (
	"net/http"
	"streetcats-api/internal/services/keycloak"
	"streetcats-api/internal/services/session"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/jinres"
	"go.uber.org/zap"
)

func cookieOptions(ctx *gin.Context) (secure bool, sameSite http.SameSite) {
	if ctx.Request.TLS != nil || strings.EqualFold(ctx.GetHeader("X-Forwarded-Proto"), "https") {
		return true, http.SameSiteNoneMode
	}
	return false, http.SameSiteLaxMode
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

type Controller struct {
	log            *zap.Logger
	kcService      keycloak.ServiceInterfaces
	sessionService session.ServiceInterfaces
	jinres         *jinres.Jinres
}

func NewAuthController(log *zap.Logger, kcService keycloak.ServiceInterfaces, sessionService session.ServiceInterfaces) *Controller {
	return &Controller{log: log, kcService: kcService, sessionService: sessionService, jinres: jinres.NewJinres()}
}

func setSessionCookie(ctx *gin.Context, sessionID string) {
	secure, sameSite := cookieOptions(ctx)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		MaxAge:   86400,
	}
	http.SetCookie(ctx.Writer, cookie)
}

func clearSessionCookie(ctx *gin.Context) {
	secure, sameSite := cookieOptions(ctx)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		MaxAge:   -1,
	}
	http.SetCookie(ctx.Writer, cookie)
}

func (c *Controller) Me(ctx *gin.Context) {
	sessionID, err := ctx.Cookie("session_id")
	if err != nil {
		c.jinres.Unauthorized().Done(ctx)
		return
	}

	session, err := c.sessionService.GetSessionByID(sessionID)
	if err != nil {
		c.jinres.Unauthorized().Done(ctx)
		return
	}

	response := map[string]interface{}{
		"username": session.Username,
	}

	c.jinres.OK().Response(response).Done(ctx)
}
