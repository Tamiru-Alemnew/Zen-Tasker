package infrastructure

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func parseToken(authHeader string) ( jwt.MapClaims, error) {
    authParts := strings.Split(authHeader, " ")

    if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
        return nil, fmt.Errorf("invalid authorization header")
    }

    token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return os.Getenv("JWT_SECRE"), nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid JWT")
    }


    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return  nil, fmt.Errorf("invalid JWT claims")
    }

    // Check token expiration
    if exp, ok := claims["exp"].(float64); ok {
        if time.Unix(int64(exp), 0).Before(time.Now()) {
            return  nil, fmt.Errorf("token has expired")
        }
    }


    return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        // JWT validation logic
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        claims , err := parseToken(authHeader)

        if err != nil {
            c.JSON(401, gin.H{"error": err.Error()})
            c.Abort()
            return
        }
        
        c.Set("claims", claims)
        c.Next()
    }
}

func RoleAuthorizationMiddleware(requiredRole string) gin.HandlerFunc{
    return func(c *gin.Context) {

        claims, exists := c.Get("claims")
        if !exists {
            c.JSON(401, gin.H{"error": "Authorization required"})
            c.Abort()
            return
        }
        role, ok := claims.(jwt.MapClaims)["role"].(string)
        if !ok {
            c.JSON(401, gin.H{"error": "Role claim is required"})
            c.Abort()
            return
        }

        if role != requiredRole {
            c.JSON(403, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
        
    }
}