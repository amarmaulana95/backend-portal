package user

type UserFormater struct { //membuat sbuah formater user
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Expired    string `json:"expired"`
}

func FormatUser(user User, token string) UserFormater { // membuat fungsi, parameternya user dan balikanya userFormater
	formatter := UserFormater{ // maping
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	return formatter
}
