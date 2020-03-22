package auth

type RegisterModel struct {
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (registerModel *RegisterModel) IsValid() bool {
	return len(registerModel.Firstname) > 0 &&
		len(registerModel.Lastname) > 0 &&
		len(registerModel.Email) > 0 &&
		len(registerModel.Password) > 0 &&
		len(registerModel.ConfirmPassword) > 0 &&
		registerModel.Password == registerModel.ConfirmPassword
}
