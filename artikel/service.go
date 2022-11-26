package article

type Service interface { //membuat interface service
	//kontrak service article
	GetArticles(userID int) ([]Article, error) // methodnya GetArticles, parameternya inputan user, balikannya user dan err
}

type service struct { //panggil repository (defidency)

	repository Repository
}

//u membuat struct service, kita butuh fungsi yg namanya NewService
func NewService(repository Repository) *service { //parameternya repository, balikannya service
	return &service{repository} // return object service yang parameternya repository
}

func (s *service) GetArticles(userID int) ([]Article, error) {

	if userID != 0 {
		//  get data atikrl berdasarkan id user
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
