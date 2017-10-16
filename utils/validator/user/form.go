package userValidator

type UserRegisterForm struct {
	Email    string `binding:"required" json:"email"`
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
