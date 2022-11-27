package article

import user "github.com/amarmaulana95/backend-portal/users"

type GetArticleDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateArticleInput struct {
	Judul             string `json:"judul" binding:"required"`
	ShortDescriptions string `json:"short_descriptions" binding:"required"`
	Descriptions      string `json:"descriptions" binding:"required"`
	User              user.User
}
