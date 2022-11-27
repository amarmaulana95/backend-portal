package handler

import (
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
func NewarticleHandler(service article.Service) *articleHandler {
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

func (h *articleHandler) GetArticle(c *gin.Context) {
	// url => api/v1/article/1
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
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", article.FormatArticle(newArticle))
	c.JSON(http.StatusOK, response)
}
