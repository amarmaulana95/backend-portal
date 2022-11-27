package article

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface { //membuat interface service
	//kontrak service article

	GetArticles(userID int) ([]Article, error)
	GetArticlesAdmin(userID int) ([]Article, error)              // methodnya GetArticles, parameternya inputan user, balikannya user dan err
	GetArticleByID(input GetArticleDetailInput) (Article, error) // methodnya GetArticleDetailInput, parameternya inputan ID, balikannya artikel dan err
	CreateArticle(input CreateArticleInput) (Article, error)
	SaveArticleImage(input CreateArticleImageInput, fileLocation string) (ArticleImages, error)
	GetArticlesUser(userID int) ([]Article, error)

	UpdateArticle(inputID GetArticleDetailInput, inputData CreateUpdateInput) (Article, error)
	// GetArticleByIDUser(input GetArticleDetailInput, userID int) (Article, error)
}

type service struct { //panggil repository (defidency)

	repository Repository
}

// UpdateArticle implements Service

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

func (s *service) GetArticlesAdmin(userID int) ([]Article, error) { //main fungsi
	if userID != 0 {
		//  get data atikel berdasarkan id user
		articles, err := s.repository.FindByUserIDAdmin(userID)
		if err != nil {
			return articles, err
		}
		return articles, nil
	}
	// else ambil smua
	articles, err := s.repository.FindAllAdmin()
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (s *service) GetArticleByID(input GetArticleDetailInput) (Article, error) { // nama fungsi
	article, err := s.repository.FindByIDU(input.ID) //cek ke repositori menggunakan parameter ID artikel
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

	newArticle, err := s.repository.Save(article) //lempar inputan ke repository
	if err != nil {
		return newArticle, err
	}

	return newArticle, nil
}

// SaveArticleImage implements Service
func (s *service) SaveArticleImage(input CreateArticleImageInput, fileLocation string) (ArticleImages, error) { // fungsi service save image
	article, err := s.repository.FindByIDU(input.ArticleID)
	if err != nil {
		return ArticleImages{}, err
	}

	if article.UserID != input.User.ID {
		return ArticleImages{}, errors.New("Not an owner of the article")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.ArticleID)
		if err != nil {
			return ArticleImages{}, err
		}
	}

	articleImage := ArticleImages{}
	articleImage.ArticleID = input.ArticleID
	articleImage.IsPrimary = isPrimary
	articleImage.FileName = fileLocation

	newAricleImage, err := s.repository.CreateImage(articleImage)
	if err != nil {
		return newAricleImage, err
	}

	return newAricleImage, nil
}

func (s *service) GetArticlesUser(userID int) ([]Article, error) { //main fungsi
	//  get data atikel berdasarkan id user
	articles, err := s.repository.FindByUserID(userID)
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (s *service) UpdateArticle(inputID GetArticleDetailInput, inputData CreateUpdateInput) (Article, error) {
	article, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return article, err
	}

	article.Approve = inputData.Approve
	article.Point = inputData.Point

	updatedArticel, err := s.repository.Update(article)
	if err != nil {
		return updatedArticel, err
	}

	return updatedArticel, nil
}
