package handler

import (
	"net/http"

	"github.com/amarmaulana95/backend-portal/helper"
	user "github.com/amarmaulana95/backend-portal/users"
	"github.com/gin-gonic/gin"
)

type userHandler struct { // handler punya deficiency terhadap service
	userService user.Service
}

//membuat fungsi handler, yg parameternya user service
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService} // karna disini akan terkoneksi dengan service
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
	formatter := user.FormatUser(NewUser, "testoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter) //panggil helper formater untuk membuat response setelah register
	c.JSON(http.StatusOK, response)                                                                    //respons 200/ok
}
