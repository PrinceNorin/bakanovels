package userValidator

type UserRegisterForm struct {
	Email           string `binding:"required,email" json:"email"`
	Username        string `binding:"required" json:"username"`
	Password        string `binding:"required" json:"password"`
	ConfirmPassword string `binding:"required,eqfield=Password" json:"confirm_password"`
}
