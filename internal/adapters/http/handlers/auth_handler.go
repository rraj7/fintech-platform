package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rraj7/demo-fintech-app/internal/service"
)

var userService = service.NewUserService()

// POST /api/auth/register
func RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		DOB      string `json:"dob"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// === Validations using UserService ===
	if !userService.ValidateUsername(req.Username) {
		c.JSON(400, gin.H{"error": "Invalid username. Must be alphanumeric, min 3 chars."})
		return
	}

	if !userService.ValidatePassword(req.Password) {
		c.JSON(400, gin.H{"error": "Password must be 8+ chars, include upper, lower, number, special char."})
		return
	}

	if !userService.ValidateEmail(req.Email) {
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	if !userService.ValidateDateOfBirth(req.DOB) {
		c.JSON(400, gin.H{"error": "DOB invalid or under 18"})
		return
	}

	if !userService.ValidatePhone(req.Phone) {
		c.JSON(400, gin.H{"error": "Invalid phone number"})
		return
	}

	if !userService.ValidateAddress(req.Address) {
		c.JSON(400, gin.H{"error": "Invalid address format"})
		return
	}

	// === Hash Password ===
	hashedPwd, err := userService.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error processing password"})
		return
	}

	// === TODO: Save user to DB ===
	// For now, mock response
	c.JSON(200, gin.H{
		"message":  "User registered successfully",
		"username": req.Username,
		"email":    req.Email,
		"password": hashedPwd, // Just for testing, NEVER return this in prod!
	})
}

// POST /api/auth/login
func LoginUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// === TODO: Fetch user & hashed password from DB ===
	// Mock hashed password for testing
	fakeHashedPwd, _ := userService.HashPassword("Test@1234")

	if !userService.CheckPasswordHash(req.Password, fakeHashedPwd) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// === TODO: Generate real JWT ===
	token := "mock-jwt-token"
	c.JSON(200, gin.H{"token": token})
}
