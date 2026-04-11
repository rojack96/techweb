package middlewares

import (
	"strings"

	"streetcats-api/configs"
	"streetcats-api/internal/services/session"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rojack96/jinres"
)

func Auth(config configs.ConfigModel, client *gocloak.GoCloak, sessionService session.ServiceInterfaces) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jr := jinres.NewJinres()

		sessionID, err := ctx.Cookie("session_id")
		if err == nil {
			session, err := sessionService.GetSessionByID(sessionID)
			if err == nil {
				ctx.Set("session", session)
				ctx.Next()
				return
			}
		}

		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			jr.Unauthorized().Done(ctx)
			ctx.Abort()
			return
		}

		authToken := strings.TrimPrefix(bearerToken, "Bearer ")

		rptResult, err := client.RetrospectToken(
			ctx,
			authToken,
			config.Keycloak.ClientId,
			config.Keycloak.ClientSecret,
			config.Keycloak.Realm,
		)
		if err != nil {
			jr.InternalServerError().Done(ctx)
			ctx.Abort()
			return
		}

		if rptResult == nil || rptResult.Active == nil || !*rptResult.Active {
			jr.Unauthorized().Done(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("authToken", authToken)

		claims, err := DecodeJWT(authToken)
		if err == nil && claims != nil {
			ctx.Set("claims", claims)
		}

		ctx.Next()
	}
}

func BasicAuth(user, pass string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jr := jinres.NewJinres()
		u, p, ok := ctx.Request.BasicAuth()

		if !ok || u != user || p != pass {
			ctx.Header("WWW-Authenticate", `Basic realm="restricted"`)

			jr.Unauthorized().Done(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// DecodeJWT - Helper per decodificare JWT senza validare firma
func DecodeJWT(authToken string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(authToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, nil
}
