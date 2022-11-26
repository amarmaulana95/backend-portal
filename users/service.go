package user

import "golang.org/x/crypto/bcrypt" //untuk password menggukan bcrypt

type Service interface { //membuat interface service
	RegisterUser(input RegisterUserInput) (User, error) // methodnya RegisterUser, parameternya inputan user, balikannya user dan err
}

type service struct {
	repository Repository //panggil repository (defidency)

}

//u membuat struct service, kita butuh fungsi yg namanya NewService
func NewService(repository Repository) *service { //parameternya repository, balikannya service
	return &service{repository} // return object service yang parameternya repository
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) { //membuat sebuah fungsi Register dengan parameter iputan struct dan membuat balikan user dan error

	user := User{}                                                                       //membuat object dari struct User
	user.Name = input.Name                                                               //isi struct
	user.Email = input.Email                                                             //isi struct
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost) //inputan pass di generate
	if err != nil {                                                                      //cek
		return user, err //return
	}
	//lanjut jika ok
	user.Password = string(password)
	user.Role = "user"
	newUser, err := s.repository.Save(user) // save user
	if err != nil {                         //cek
		return newUser, err //jika err
	}
	//lanjut ok
	return newUser, nil

}

// maping struct inputan -> struct User
// simpan struct User melalui repository
