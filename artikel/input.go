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

type CreateUpdateInput struct {
	Approve int `json:"approve" binding:"required"`
	Point   int `json:"point" binding:"required"`
	User    user.User
}

type CreateArticleImageInput struct {
	ArticleID int  `form:"article_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}
