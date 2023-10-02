package utils

import (
	"OIDC/model/response"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

const jwt_secret = "bAlc5pLZek78sOuVZm0p6L3OmY1qSIb8u3ql"

// GenerateToken generates a jwt token
func GenerateJWTToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	// 这里不会有error被返回
	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		log.Println("[error]:", err)
		return ""
	}
	return tokenString
}

// ParseToken parses a jwt token

func MiddlewareJWTAuthorize() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if len(auth) <= len("Bearer ") {
			c.JSON(http.StatusUnauthorized, response.UnauthorizedError)
			c.Abort()
			return
		}
		tokenString := auth[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_secret), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.UnauthorizedError)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.UnauthorizedError)
			c.Abort()
			return
		}
		id, err := strconv.ParseUint((*claims)["userID"].(string), 10, 32)
		if err != nil {
			log.Println("[error]:", err)
			c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("校验令牌失败"))
			c.Abort()
			return
		}
		c.Set("userID", uint(id))
	}
}
