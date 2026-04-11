package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	state := generateRandomState()

	// salva lo state in cookie per verificare al ritorno
	setStateCookie(ctx, state)

	url := c.kcService.GetLoginURL(state)

	ctx.Redirect(302, url)
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
