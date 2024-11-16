package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"task-golang-db/models" // Import the correct path to the models package

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("your_secret_key")

// VerifyToken middleware to authenticate user and return user info from token
func VerifyToken(c *gin.Context) (models.Auth, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return models.Auth{}, errors.New("authorization header required")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return models.Auth{}, errors.New("invalid or expired token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Retrieve user details from token claims
		user := models.Auth{
			AuthID:    int64(claims["auth_id"].(float64)),
			AccountID: int64(claims["account_id"].(float64)),
			Username:  claims["username"].(string),
		}
		return user, nil
	}

	return models.Auth{}, errors.New("failed to parse token claims")
}

// AuthMiddleware validates JWT token and adds claims to context
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Printf("Token claims: %v", claims)
			if authID, ok := claims["auth_id"].(float64); ok {
				c.Set("auth_id", int64(authID))
			}
			if accountID, ok := claims["account_id"].(float64); ok {
				c.Set("account_id", int64(accountID))
			}
			if username, ok := claims["username"].(string); ok {
				c.Set("username", username)
			}
		} else {
			log.Printf("Invalid claims: %v", token.Claims)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
