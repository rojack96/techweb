package auth

import (
	"streetcats-api/internal/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Login(ctx *gin.Context) {
	var req dto.AuthDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.jinres.BadRequest().Message("invalid_request").Done(ctx)
		return
	}

	token, err := c.kcService.LoginDirect(req.Username, req.Password)
	if err != nil {
		c.log.Error("login failed", zap.Error(err))
		c.jinres.Unauthorized().Message("invalid_credentials").Done(ctx)
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
	c.jinres.OK().Custom("username", derefString(userInfo.PreferredUsername)).Message("login successful").Done(ctx)
}
