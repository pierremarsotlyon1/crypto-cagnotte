package auth

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginModel *LoginModel) IsValid() bool {
	return len(loginModel.Email) > 0 && len(loginModel.Password) > 0
}
