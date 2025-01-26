package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        fmt.Println("Token received:", tokenString) // Log token yang diterima

        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(jwtKey), nil
        })

        if err != nil || !token.Valid {
            fmt.Println("Token parsing error:", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || claims["role"] != "Admin" {
            fmt.Println("Unauthorized claims:", claims)
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin access only"})
            c.Abort()
            return
        }

        // Set userID dan role ke context
        c.Set("userID", claims["userID"])
        c.Set("role", claims["role"])
        c.Next()
    }
}


