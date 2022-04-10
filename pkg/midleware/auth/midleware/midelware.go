package midleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"users-api-server/pkg/midleware/auth"
)

type AuthMiddleware struct {
	authSrv auth.IAuthService
}

func NewAuthMiddleware(authSrv auth.IAuthService) gin.HandlerFunc {
	return (&AuthMiddleware{
		authSrv: authSrv,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if _, err := m.authSrv.ValidateToken(headerParts[1]); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
