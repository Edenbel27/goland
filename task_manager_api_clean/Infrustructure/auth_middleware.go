package Infrustructure


import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tok := c.GetHeader("Authorization")
		if tok == "" || !strings.HasPrefix(tok, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(tok, "Bearer"))
		token , err := jwt.Parse(tokenStr, func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil , fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRETKEY")) , nil
		})

		
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Claims"})
			c.Abort()
			return

		}
		c.Set("userEmail", claims["email"].(string))
		c.Set("userRole", claims["role"].(string))
		c.Next()
	}
}

func AuthorizeRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "does not exist"})
			c.Abort()
			return
		}
		if userRole != role{
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden", "role":userRole})
			c.Abort()
			return
		}
		c.Next()
	}
}
