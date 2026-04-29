package user

import (
	"time"

	"coffee-pos-api/internal/config"
	"coffee-pos-api/internal/database"
	"coffee-pos-api/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	var user User

	if err := database.DB.
		Where("email = ?", req.Email).
		First(&user).Error; err != nil {

		response.Error(c, 401, "Invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {

		response.Error(c, 401, "Invalid credentials")
		return
	}

	cfg := config.LoadConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		response.Error(c, 500, "Failed to generate token")
		return
	}

	response.Success(c, 200, "Login successful", LoginResponse{
		Token: tokenString,
	})
}

func Profile(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var user User

	if err := database.DB.First(&user, userID).Error; err != nil {
		response.Error(c, 404, "User not found")
		return
	}

	response.Success(c, 200, "Profile fetched", ToProfileResponse(user))
}
