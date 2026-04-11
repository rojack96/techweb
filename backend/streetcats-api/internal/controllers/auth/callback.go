package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) CallbackHandler(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")

	if code == "" {
		c.jinres.BadRequest().Message("missing_code").Done(ctx)
		return
	}

	if state == "" {
		c.jinres.BadRequest().Message("missing_state").Done(ctx)
		return
	}

	// verifica lo state dal cookie
	cookieState, err := ctx.Cookie("oauth_state")
	if err != nil || cookieState != state {
		c.log.Error("state mismatch or missing", zap.String("received_state", state), zap.String("cookie_state", cookieState))
		c.jinres.BadRequest().Message("invalid_state").Done(ctx)
		return
	}

	// cancella il cookie dello state dopo la verifica
	clearStateCookie(ctx)

	token, err := c.kcService.ExchangeCode(code)
	if err != nil {
		c.log.Error("exchange code failed", zap.Error(err))
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	accessToken := token.AccessToken
	if accessToken == "" {
		c.log.Error("empty access token returned from Keycloak")
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	userInfo, err := c.kcService.GetUserInfo(accessToken)
	if err != nil {
		c.log.Error("user info retrieval failed", zap.Error(err))
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	sessionID, err := c.sessionService.Create(userInfo, token)
	if err != nil {
		c.log.Error("session creation failed", zap.Error(err))
		c.jinres.InternalServerError().Done(ctx)
		return
	}

	setSessionCookie(ctx, sessionID)
	ctx.Redirect(http.StatusFound, "/")
}
