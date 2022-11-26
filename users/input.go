package user

type RegisterUserInput struct { //membuat struct inputan register, yang mewakili inputan user/nanti di postman
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct { //struct inputan login
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
