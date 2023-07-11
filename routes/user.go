package routes

import (
	"github.com/gin-gonic/gin"
	"golang_basic_gin/auth"
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	var user models.User

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}

	//hash password
	err = user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	//insert user to DB
	insert := db.Create(&user)
	if insert.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":  "Insert Sukses",
		"user id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

func GenereteToken(c *gin.Context) {

	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	requets := models.TokenRequest{}
	user := models.User{}

	err = c.ShouldBindJSON(&requets)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}

	//check email
	checkEmail := db.Where("email = ?", requets.Email).First(&user)
	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found"})
		return
	}
	//check password
	credentiaError := user.ChechPassword(requets.Password)
	if credentiaError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password Not Match"})
		return
	}

	//generte token
	tokenString, err := auth.GenereteJWT(user.Email, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Generete JWT"})
		return
	}

	//resposnes

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
	})
}
