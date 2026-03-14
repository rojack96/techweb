package middlewares

import (
	"strings"

	"streetcats-api/configs"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rojack96/jinres"
)

func Auth(config configs.ConfigModel, client *gocloak.GoCloak, prefixPath string) gin.HandlerFunc {
	// This is a workaround to disable a Keycloak auth on a specific path
	// in that case is a path with another auth mode (basic auth)
	skipPaths := []string{
		"user/register",
	}

	return func(ctx *gin.Context) {
		path := ctx.FullPath()

		for _, skip := range skipPaths {
			if path == prefixPath+skip {
				ctx.Next()
				return
			}
		}

		jr := jinres.NewJinres()
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
