package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(user RegisterUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUsers() ([]GetUserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func hashPassword(input string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}

	return string(passwordHash)
}

func dateFormater(input string) string {
	tm, err := time.Parse("2006-01-02", input)
	if err != nil {
		return err.Error()
	}

	return string(tm.Format("2006-01-02"))
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	dataUser := User{
		Name:        input.Name,
		Gender:      input.Gender,
		Dateofbirth: dateFormater(input.Dateofbirth),
		Email:       input.Email,
		Password:    hashPassword(input.Password),
	}

	newUser, err := s.repository.AddUser(dataUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUsers() ([]GetUserFormat, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var dataUser GetUserFormat
	var dataUserList []GetUserFormat
	for _, u2 := range users {
		dataUser = GetUserFormat{
			ID:    u2.ID,
			Name:  u2.Name,
			Email: u2.Email,
		}
		dataUserList = append(dataUserList, dataUser)
	}

	return dataUserList, nil
}

// func (s *service) copyEntity(data []User) []GetUserFormat {
// 	models := make([]GetUserFormat, len(data))
// 	for i, v := range data {
// 		models[i].Name = v.Name
// 		models[i].Gender = v.Gender
// 		models[i].Email = v.Email
// 	}
// 	return models
// }
