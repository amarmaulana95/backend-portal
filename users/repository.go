package user

import "gorm.io/gorm"

type Repository interface { // interface repository,
	//construct

	FindAll() ([]User, error)
	Save(user User) (User, error)           //membuat sebuah fungsi save yang parameternya struc User, dan mengembalikan user/err
	FindByEmail(email string) (User, error) //membuat sebuah fungsi FindByEmail yang parameternya string, dan mengembalikan user/err
	FindByID(ID int) (User, error)          //membuat sebuah fungsi FindByID yang parameternya ID, retunr User dan err
}

type repository struct { // sebuah struct bernama repository (r nya kecil) yang artinya tidak bersifat public/tidak bisa di panggil di package yg lain.
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository { //membuat sbuah object baru dari struct repository
	return &repository{db} //isi nilai dari db menggunakan parameter , (maaf agak susah menjelaskan tp sy paham)
}

func (r *repository) Save(user User) (User, error) { //membuat sebuah fungsi yang bernama Save mengikuti signature di atas/interface dengan mengambil parameter
	err := r.db.Debug().Create(&user).Error // untuk save inputan ke db dengan pointer &
	if err != nil {                         //cek ada error ga ?
		return user, err //return user dan error
	}

	return user, nil //return jika berhasil

}

func (r *repository) FindByEmail(email string) (User, error) { // fungsi find by email dengan parameter email, mengembalikan user dan err
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error //cek ke db ada user dengan email ...?

	if err != nil { //cek validasi
		return user, nil //return
	}
	return user, nil //jika ada balikan user dan nil
}

func (r *repository) FindByID(ID int) (User, error) { // fungsi find by ID dengan parameter email
	var user User
	// object user
	err := r.db.Where("id = ?", ID).Find(&user).Error
	//cek di db dengan ID
	if err != nil {
		//validasi
		return user, nil
	}
	//lanjut return ok
	return user, nil
}

func (r *repository) FindAll() ([]User, error) { // fungsi get all user
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
