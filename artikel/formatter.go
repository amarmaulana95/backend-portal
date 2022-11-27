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
	//array article formater
	articlesFormatter := []ArticleFormatter{} //jika data kosong balikan ke array kosong

	for _, article := range articles {
		articleFormatter := FormatArticle(article)
		articlesFormatter = append(articlesFormatter, articleFormatter)
	}
	return articlesFormatter
}

type ArticleDetailFormatter struct { // membuat formater untuk kebutuhan reponse detail
	ID                int                     `json:"id"`
	Name              string                  `json:"name"`
	ShortDescriptions string                  `json:"short_descriptions"`
	Description       string                  `json:"description"`
	ImageURL          string                  `json:"image_url"`
	UserID            int                     `json:"user_id"`
	Slug              string                  `json:"slug"`
	User              ArticleUserFormatter    `json:"user"`
	Images            []ArticleImageFormatter `json:"images"`
}

type ArticleUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ArticleImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatArticleDetail(article Article) ArticleDetailFormatter { //=> membuat formater untuk kebutuhan reponse detail
	articleDetailFormatter := ArticleDetailFormatter{}
	articleDetailFormatter.ID = article.ID
	articleDetailFormatter.Name = article.Judul
	articleDetailFormatter.ShortDescriptions = article.ShortDescriptions
	articleDetailFormatter.Description = article.Description
	articleDetailFormatter.UserID = article.UserID
	articleDetailFormatter.Slug = article.Slug
	articleDetailFormatter.ImageURL = ""

	if len(article.ArticleImages) > 0 {
		articleDetailFormatter.ImageURL = article.ArticleImages[0].FileName
	}

	user := article.User

	articleUserFormatter := ArticleUserFormatter{}
	articleUserFormatter.Name = user.Name
	articleUserFormatter.ImageURL = user.Avatar

	articleDetailFormatter.User = articleUserFormatter

	images := []ArticleImageFormatter{}

	for _, image := range article.ArticleImages {
		articleImageFormatter := ArticleImageFormatter{}
		articleImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		articleImageFormatter.IsPrimary = isPrimary

		images = append(images, articleImageFormatter)
	}

	articleDetailFormatter.Images = images

	return articleDetailFormatter
}
