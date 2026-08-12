// package handler

// import (
// 	"encoding/json"
// 	"net/http"

// 	"task-management/models"

// 	"task-management/utils"

// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// type RegisterRequest struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // =========================
// // REGISTER
// // =========================

// func Register(w http.ResponseWriter, r *http.Request) {

// 	var req RegisterRequest

// 	// Request body se JSON lena
// 	err := json.NewDecoder(r.Body).Decode(&req)

// 	if err != nil {
// 		http.Error(
// 			w,
// 			"Something went wrong!",
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}

// 	// Password ko hash karna
// 	hashedPassword, err := bcrypt.GenerateFromPassword(
// 		[]byte(req.Password),
// 		bcrypt.DefaultCost,
// 	)

// 	if err != nil {
// 		http.Error(
// 			w,
// 			"Cannot hash password!",
// 			http.StatusInternalServerError,
// 		)
// 		return
// 	}

// 	// User banana
// 	user := models.User{
// 		Name:     req.Name,
// 		Email:    req.Email,
// 		Password: string(hashedPassword),
// 	}

// 	// Database mein user save karna
// 	if err := DB.Create(&user).Error; err != nil {
// 		http.Error(
// 			w,
// 			"Failed to create user",
// 			http.StatusInternalServerError,
// 		)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "User successfully created",
// 	})
// }

// // =========================
// // LOGIN
// // =========================

// func Login(w http.ResponseWriter, r *http.Request) {

// 	var req LoginRequest

// 	// Request body read
// 	err := json.NewDecoder(r.Body).Decode(&req)

// 	if err != nil {
// 		http.Error(
// 			w,
// 			"Invalid request",
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}

// 	// Database se user find karo
// 	var user models.User

// 	result := DB.
// 		Where("email = ?", req.Email).
// 		First(&user)

// 	if result.Error != nil {
// 		http.Error(
// 			w,
// 			"Invalid email or password",
// 			http.StatusUnauthorized,
// 		)
// 		return
// 	}

// 	// User ka entered password
// 	// database wale hashed password se compare karo
// 	err = bcrypt.CompareHashAndPassword(
// 		[]byte(user.Password),
// 		[]byte(req.Password),
// 	)

// 	if err != nil {
// 		http.Error(
// 			w,
// 			"Invalid email or password",
// 			http.StatusUnauthorized,
// 		)
// 		return
// 	}

// 	// Password correct hai
// 	// Ab JWT token generate karo
// 	token, err := utils.GenerateToken(user.ID)

// 	if err != nil {
// 		http.Error(
// 			w,
// 			"Could not generate token",
// 			http.StatusInternalServerError,
// 		)
// 		return
// 	}

// 	// JWT frontend ko bhejo
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Login successful",
// 		"token":   token,
// 	})
// }

package controllers

import (
	"net/http"

	"task-management/database"
	"task-management/models"
	"task-management/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// =========================
// REQUEST STRUCTS
// =========================

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// =========================
// REGISTER
// =========================

func Register(c *gin.Context) {

	var req RegisterRequest

	// Request body se JSON lena
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Password hash karna
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot hash password",
		})
		return
	}

	// User model banana
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Database mein user save karna
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created",
	})
}

// =========================
// LOGIN
// =========================

func Login(c *gin.Context) {

	var req LoginRequest

	// Request body se JSON lena
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Database se user find karna
	var user models.User

	result := database.DB.
		Where("email = ?", req.Email).
		First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Password check
	// req.Password = user ne jo password diya
	// user.Password = database mein stored hashed password
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Password correct hai
	// JWT token generate karo
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token",
		})
		return
	}

	// Login response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",

		"token": token,

		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
