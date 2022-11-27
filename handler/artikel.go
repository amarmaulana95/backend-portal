package handler

import (
	"fmt"
	"net/http"
	"strconv"

	article "github.com/amarmaulana95/backend-portal/artikel"
	"github.com/amarmaulana95/backend-portal/helper"
	user "github.com/amarmaulana95/backend-portal/users"
	"github.com/gin-gonic/gin"
)

// tangkap param di handler
// handler ke service
// service menentukan repository mana yg akan di call
// repo :  panggil fungsi ...
// to db

type articleHandler struct {
	service article.Service
}

//membuat fungsi handler, yg parameternya apapun berkaitan dengan service
func NewArticleHandler(service article.Service) *articleHandler {
	return &articleHandler{service} // karna disini akan terkoneksi dengan service
}

//api/v1/article
func (h *articleHandler) GetArticles(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id")) //ubah ke int dngn strconv.atoi

	articles, err := h.service.GetArticles(userID)
	if err != nil {
		response := helper.APIResponse("Error get data posted ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// response := helper.APIResponse("List data posted all", http.StatusOK, "success", articles)
	response := helper.APIResponse("List data posted all", http.StatusOK, "success", article.FormatArticles(articles))
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetArticlesAdmin(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id")) //ubah ke int dngn strconv.atoi

	currentUser := c.MustGet("currentUser").(user.User)
	UserRole := currentUser.Role

	if UserRole != "admin" { //cek yg update admin bukan ? validasi tolak jika bukan admin
		response := helper.APIResponse("error", http.StatusUnprocessableEntity, "error", "you not admin")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	articles, err := h.service.GetArticlesAdmin(userID)
	if err != nil {
		response := helper.APIResponse("Error get data posted ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// response := helper.APIResponse("List data posted all", http.StatusOK, "success", articles)
	response := helper.APIResponse("List data posted all", http.StatusOK, "success", article.FormatArticles(articles))
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetArticle(c *gin.Context) {
	// => handler : maping inputan id di url ke struct input => service, call formatter
	var input article.GetArticleDetailInput

	err := c.ShouldBindUri(&input) //maping inputan
	if err != nil {                //cek ada eror ?
		response := helper.APIResponse("Failed to get detail of article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	articleDetail, err := h.service.GetArticleByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// jika ok masukan balikan respon ke formater lalu show
	response := helper.APIResponse("Article detail", http.StatusOK, "success", article.FormatArticleDetail(articleDetail))
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) CreateArticle(c *gin.Context) {
	var input article.CreateArticleInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User) //untuk mendaptkan data curent user

	input.User = currentUser

	newArticle, err := h.service.CreateArticle(input)
	if err != nil {
		response := helper.APIResponse("Failed to create article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create article", http.StatusOK, "success", article.FormatArticle(newArticle))
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) UploadImage(c *gin.Context) { //fungsi upload image
	var input article.CreateArticleImageInput //maping inputan

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload article image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload article image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload article image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveArticleImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload article image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("article image successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetArticleUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	articles, err := h.service.GetArticlesUser(userID)
	if err != nil {
		response := helper.APIResponse("Error get data posted ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// response := helper.APIResponse("List data posted all", http.StatusOK, "success", articles)
	response := helper.APIResponse("List data posted all", http.StatusOK, "success", article.FormatArticles(articles))
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) UpdateArticle(c *gin.Context) {
	var inputID article.GetArticleDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData article.CreateUpdateInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	UserRole := currentUser.Role

	if UserRole != "admin" { //cek yg update admin bukan ? validasi tolak jika bukan admin
		response := helper.APIResponse("error", http.StatusUnprocessableEntity, "error", "you not admin")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//jika admin approve
	result, err := h.service.UpdateArticle(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update article", http.StatusOK, "success", article.FormatArticle(result))
	c.JSON(http.StatusOK, response)
}
