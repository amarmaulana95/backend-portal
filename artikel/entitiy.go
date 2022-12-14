package article

import (
	"time"

	user "github.com/amarmaulana95/backend-portal/users"
)

type Article struct {
	ID                int
	UserID            int
	Judul             string
	ShortDescriptions string
	Descriptions      string
	Slug              string
	Point             int
	Approve           int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArticleImages     []ArticleImages
	User              user.User
}

type ArticleImages struct {
	ID        int
	ArticleID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
