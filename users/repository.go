package user

import "gorm.io/gorm"

type Repository interface { //membuat sebuah interface yg bernama repository,

	Save(user User) (User, error) //membuat sebuah fungsi save yang parameternya struc User, dan mengembalikan user/err
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
