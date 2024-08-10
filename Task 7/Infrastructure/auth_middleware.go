package infrastructure

import (
	"strings"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func JWTValidation() gin.HandlerFunc{
	return func(c *gin.Context){
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		
		authSlice := strings.Split(authHeader, " ")
		if len(authSlice) != 2 || strings.ToLower(authSlice[0]) != "bearer"{
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := ValidateToken(authSlice[1])
		  
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("claims", claims)
		}else{
			c.JSON(401, gin.H{"error": "invalid auth claims"})
			c.Abort()
			return 
		}
		c.Next()
	}
}

func RoleAuth(roles ...string) gin.HandlerFunc{
	return func (c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(401, gin.H{
				"error":"claims not found",
			})
			c.Abort()
			return 
		}
		claimsHash, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		role, ok := claimsHash["role"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Role not found in JWT claims"})
			c.Abort()
			return
		}

		role_authorized := false 
		for _, elem := range roles{
			if elem == role{
				role_authorized = true
				break
			}
		}
		if !role_authorized{
			c.JSON(401, gin.H{"error": "Your role doesn't have access to this resource"})
			c.Abort()
			return
		}
		c.Set("userID", claimsHash["user_id"])
		c.Next()
	}
}