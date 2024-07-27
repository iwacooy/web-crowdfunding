package user

type UserFormat struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Pekerjaan string `json:"pekerjaan"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func NewUserFormat(user User) UserFormat {
	formatUser := UserFormat{
		ID:        user.ID,
		Name:      user.Nama,
		Pekerjaan: user.Pekerjaan,
		Email:     user.Email,
		Token:     user.Token,
	}

	return formatUser
}
