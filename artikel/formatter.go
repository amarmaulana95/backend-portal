package article

type ArticleFormatter struct {
	ID                int    `json:"id"`
	userID            int    `json:"user_id"`
	Judul             string `json:"judul"`
	ShortDescriptions string `json:"short_descriptions"`
	ImageUrl          string `json:"image_url"`
	Slug              string `json:"slug"`
}

func FormatArticle(article Article) ArticleFormatter {
	//buat object
	ArticleFormatter := ArticleFormatter{}
	ArticleFormatter.ID = article.ID
	ArticleFormatter.userID = article.UserID
	ArticleFormatter.Judul = article.Judul
	ArticleFormatter.ShortDescriptions = article.ShortDescriptions
	ArticleFormatter.Slug = article.Slug
	ArticleFormatter.ImageUrl = ""

	if len(article.ArticleImages) > 0 {
		ArticleFormatter.ImageUrl = article.ArticleImages[0].FileName
	}
	return ArticleFormatter
}

func FormatArticles(articles []Article) []ArticleFormatter {
	//array campaign formater
	articlesFormatter := []ArticleFormatter{} //jika data kosong balikan ke array kosong

	for _, article := range articles {
		articleFormatter := FormatArticle(article)
		articlesFormatter = append(articlesFormatter, articleFormatter)
	}
	return articlesFormatter
}
