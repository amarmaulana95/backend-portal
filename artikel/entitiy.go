package article

import (
	"os/user"
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
	User             user.User
}

type ArticleImages struct {
	ID        int
	ArticleID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
