package user

type RegisterUserInput struct {
	Nama      string `json:"nama" binding:"required"`
	Pekerjaan string `json:"pekerjaan" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Token     string `json:"token"`
}
