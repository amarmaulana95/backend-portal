package article

type GetArticleDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
