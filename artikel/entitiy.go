package article

import (
	"time"
)

type Article struct {
	ID               int
	UserID           int
	Judul            string
	ShortDescription string
	Description      string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ArticleImages    []ArticleImages
}

type ArticleImages struct {
	ID        int
	ArticleID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
