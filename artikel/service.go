package article

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface { //membuat interface service
	//kontrak service article
	GetArticles(userID int) ([]Article, error)                   // methodnya GetArticles, parameternya inputan user, balikannya user dan err
	GetArticleByID(input GetArticleDetailInput) (Article, error) // methodnya GetArticleDetailInput, parameternya inputan ID, balikannya artikel dan err
	CreateArticle(input CreateArticleInput) (Article, error)
}

type service struct { //panggil repository (defidency)

	repository Repository
}

//u membuat struct service, kita butuh fungsi yg namanya NewService
func NewService(repository Repository) *service { //parameternya repository, balikannya service
	return &service{repository} // return object service yang parameternya repository
}

func (s *service) GetArticles(userID int) ([]Article, error) { //main fungsi
	if userID != 0 {
		//  get data atikel berdasarkan id user
		articles, err := s.repository.FindByUserID(userID)
		if err != nil {
			return articles, err
		}
		return articles, nil
	}
	// else ambil smua
	articles, err := s.repository.FindAll()
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (s *service) GetArticleByID(input GetArticleDetailInput) (Article, error) { // nama fungsi
	article, err := s.repository.FindByID(input.ID) //cek ke repositori menggunakan parameter ID artikel
	if err != nil {
		return article, err
	}
	return article, nil
}

func (s *service) CreateArticle(input CreateArticleInput) (Article, error) { //fungsi create artikel
	//maping inputan ke object article
	article := Article{} // => object
	article.Judul = input.Judul
	article.ShortDescriptions = input.ShortDescriptions
	article.Descriptions = input.Descriptions
	article.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Judul, input.User.ID) //membuat slug dengan library slug
	article.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(article) //lempar inputan ke repository
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
