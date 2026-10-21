package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MGT06/shortlink-app_backend/internal/dto"
	"github.com/MGT06/shortlink-app_backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func CheckToken(rc *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "please login first",
			})
			return
		}

		result := strings.Split(bearer, " ")
		if len(result) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}
		if result[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}

		var token pkg.JWTClaims
		err := token.DecodeToken(result[1])
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
					Success: false,
					Message: "invalid token",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Message: "a system error has occurred",
			})
			return
		}

		key := fmt.Sprintf("eventhub:tokenBlacklist:%s", token.ID)
		n, err := rc.Exists(c.Request.Context(), key).Result()
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Message: "a system error has occurred",
			})
			return
		}
		if n > 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "invalid token",
			})
			return
		}

		c.Set("userId", token.Id)
		c.Set("jti", token.ID)
		c.Set("exp", token.ExpiresAt.Time)
		c.Next()
	}
}
