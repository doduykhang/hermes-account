package service

import (
	"doduykhang/hermes-account/pkg/dto"
	"doduykhang/hermes-account/pkg/model"
	"doduykhang/hermes-account/pkg/repository"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account interface {
	Register(request dto.RegisterRequest) (string, error)
	Login(request dto.LoginRequest) (string, error)
}

type account struct {
	repo repository.Account
}

func NewAccount(repo repository.Account) Account {
	return &account{
		repo: repo,
	}
}

func (a *account) Register(request dto.RegisterRequest) (string, error) {
	var account model.Account
	account.UserID = uuid.New().String()
	account.ID = uuid.New().String()
	account.Email = request.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", nil
	}
	account.Password = string(hashedPassword)
	_, err = a.repo.CreateAccount(&account)
	if err != nil {
		return "", err
	}	
	return account.UserID, nil
}
func (a *account) Login(request dto.LoginRequest) (string, error) {
	account, err := a.repo.FindAccountByEmail(request.Email)
	if err != nil {
		return "", nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("Wrong user name or password")
	}
	return account.UserID, nil
}


