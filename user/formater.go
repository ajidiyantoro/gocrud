package user

type UserFormater struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}

type GetUserFormat struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Test []struct{}

func FormatUser(user User, token string) UserFormater {
	formater := UserFormater{
		ID:          user.ID,
		Name:        user.Name,
		Gender:      user.Gender,
		Dateofbirth: user.Dateofbirth,
		Email:       user.Email,
		Token:       token,
	}

	return formater
}
