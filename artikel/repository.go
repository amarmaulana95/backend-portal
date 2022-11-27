package article

import "gorm.io/gorm"

type Repository interface { // seperti biasa membuat interface repository
	//di definisikan dengan nama ...
	FindAll() ([]Article, error)
	FindByUserID(userID int) ([]Article, error)
	FindByID(ID int) (Article, error)
	Save(article Article) (Article, error)
}

type repository struct { // sebuah struct bernama repository (r nya kecil) yang artinya tidak bersifat public/tidak bisa di panggil di package yg lain.
	db *gorm.DB
}

func (Article) TableName() string {
	return "article"
}

func NewRepository(db *gorm.DB) *repository { //membuat sbuah object baru dari struct repository
	return &repository{db} //isi nilai dari db menggunakan parameter , (maaf agak susah menjelaskan tp sy paham)
}

func (r *repository) FindAll() ([]Article, error) { //membuat fungsi untuk mencari smua artikel db
	var articles []Article

	err := r.db.Preload("ArticleImages", "article_images.is_primary = 1").Find(&articles).Error
	if err != nil { //validasi

		return articles, err //return jika err
	}
	return articles, nil //return jika ok
}

func (r *repository) FindByUserID(userID int) ([]Article, error) { //membuat fungsi untuk mencari  artikel berdasarkan userID
	var articles []Article
	err := r.db.Where("user_id = ?", userID).Preload("ArticleImages", "article_images.is_primary = 1").Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (r *repository) FindByID(ID int) (Article, error) { //membuat fungsi untuk mencari artikel berdasarkan artikel ID
	var article Article

	err := r.db.Preload("User").Preload("ArticleImages").Where("id = ?", ID).Find(&article).Error

	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) Save(article Article) (Article, error) { //membuatfungsi untuk save dengan parameter inputan postman yg sudah di maping berdasarkan struct
	err := r.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}
