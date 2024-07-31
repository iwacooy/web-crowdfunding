package user

type UserFormat struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Pekerjaan string `json:"pekerjaan"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func NewUserFormat(user User, token string) UserFormat {
	formatUser := UserFormat{
		ID:        user.ID,
		Nama:      user.Nama,
		Pekerjaan: user.Pekerjaan,
		Email:     user.Email,
		Token:     token,
	}

	return formatUser
}
