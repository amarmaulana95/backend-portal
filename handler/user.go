package handler

import (
	"net/http"

	"github.com/amarmaulana95/backend-portal/auth"
	"github.com/amarmaulana95/backend-portal/helper"
	user "github.com/amarmaulana95/backend-portal/users"
	"github.com/gin-gonic/gin"
)

type userHandler struct { // handler punya deficiency terhadap service
	//definisi
	userService user.Service
	authService auth.Service
}

//membuat fungsi handler, yg parameternya apapun berkaitan dengan service
func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService} // karna disini akan terkoneksi dengan service
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap inputan user/ aku mendapatka mu wahai inputan user
	// map input dari user ke struct RegisrterUserInput
	// struct diatas akan di parsing sbg param service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil { // cek inputan
		errors := helper.FormatValidationError(err) //get metode helper error
		errorMessage := gin.H{"errors": errors}     // error tsb di map oleh gin.H
		response := helper.APIResponse("failed register", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// lanjut
	NewUser, err := h.userService.RegisterUser(input)
	if err != nil { //cek inputan
		response := helper.APIResponse("Register failed", http.StatusBadRequest, "error", nil) //panggil helper formater untuk response error
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//jika ok
	token, err := h.authService.GenerateToken(NewUser.ID, NewUser.Role)
	if err != nil {
		response := helper.APIResponse("failed register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(NewUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter) //panggil helper formater untuk membuat response setelah register
	c.JSON(http.StatusOK, response)                                                                    //respons 200/ok
}

func (h *userHandler) Login(c *gin.Context) {
	// user input email dan pass
	// input ditangkap handler
	// maping inputan ke input struct
	// input strcuk parsing ke service
	// service mencari dg bantuan repository user dengan email
	// if ketemu cocokan password

	var input user.LoginInput // tangkap inputan login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err) //get metode helper error
		errorMessage := gin.H{"errors": errors}     // error tsb di map oleh gin.H

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) //return err jika gagal
		return
	}
	//Lanjut validasi
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID, loggedinUser.Role) //generate token
	if err != nil {
		response := helper.APIResponse("Loggin failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//jika OK
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	UserRole := currentUser.Role
	if UserRole != "admin" { //cek yg update admin bukan ? validasi tolak jika bukan admin
		response := helper.APIResponse("error", http.StatusUnprocessableEntity, "error", "you not admin")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, err := h.userService.GetAllUsers()
	if err != nil {
		response := helper.APIResponse("Error get data user ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// response := helper.APIResponse("List data posted all", http.StatusOK, "success", articles)
	response := helper.APIResponse("List data user all", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, response)

}
