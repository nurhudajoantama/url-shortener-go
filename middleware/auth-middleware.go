package middleware

import (
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	jwtService  service.JwtService
	authService service.AuthService
}

type AuthMiddleware interface {
	JwtAuthMiddleware(*gin.Context)
}

func NewAuthMiddleware(as service.AuthService, js service.JwtService) AuthMiddleware {
	return &authMiddleware{
		authService: as,
		jwtService:  js,
	}
}

func (m *authMiddleware) JwtAuthMiddleware(c *gin.Context) {
	// Get the client token
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatus(401)
		return
	}

	claims, err := m.jwtService.GetClaimsByToken(token)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	username := claims["username"]
	user, err := m.authService.FindByUsername(username.(string))
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	// Check TIME !!!

	c.Set("user", user)
	c.Next()
}
