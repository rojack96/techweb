package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"streetcats-api/internal/services/keycloak"
	"streetcats-api/internal/services/session"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/jinres"
	"go.uber.org/zap"
)

type Controller struct {
	log            *zap.Logger
	kcService      keycloak.ServiceInterfaces
	sessionService session.ServiceInterfaces
	jinres         *jinres.Jinres
}

func NewAuthController(log *zap.Logger, kcService keycloak.ServiceInterfaces, sessionService session.ServiceInterfaces) *Controller {
	return &Controller{log: log, kcService: kcService, sessionService: sessionService, jinres: jinres.NewJinres()}
}

func generateRandomState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func setStateCookie(ctx *gin.Context, state string) {
	cookie := &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // set true in production with HTTPS
		SameSite: http.SameSiteNoneMode,
		MaxAge:   300, // 5 minuti
	}
	http.SetCookie(ctx.Writer, cookie)
}

func setSessionCookie(ctx *gin.Context, sessionID string) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // set true in production with HTTPS
		SameSite: http.SameSiteNoneMode,
		MaxAge:   86400,
	}
	http.SetCookie(ctx.Writer, cookie)
}

func clearSessionCookie(ctx *gin.Context) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1,
	}
	http.SetCookie(ctx.Writer, cookie)
}

func clearStateCookie(ctx *gin.Context) {
	cookie := &http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1,
	}
	http.SetCookie(ctx.Writer, cookie)
}
