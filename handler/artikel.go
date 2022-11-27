package handler

import (
	"net/http"
	"strconv"

	article "github.com/amarmaulana95/backend-portal/artikel"
	"github.com/amarmaulana95/backend-portal/helper"
	"github.com/gin-gonic/gin"
)

// tangkap param di handler
// handler ke service
// service menentukan repository mana yg akan di call
// repo : GetAll GetUserbyID
// to db

type articleHandler struct {
	service article.Service
}

//membuat fungsi handler, yg parameternya apapun berkaitan dengan service
func NewarticleHandler(service article.Service) *articleHandler {
	return &articleHandler{service} // karna disini akan terkoneksi dengan service
}

//api/v1/campaigns
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
