package repository

import (
	"doduykhang/hermes-account/pkg/model"

	"gorm.io/gorm"
)

type Account interface {
	CreateAccount(account *model.Account) (*model.Account, error)
	FindAccountByEmail(email string) (*model.Account, error)
}

type account struct {
	db *gorm.DB	
}

func NewAccount(db *gorm.DB) Account {
	return &account{
		db: db,
	}
}

func (r *account) CreateAccount(account *model.Account) (*model.Account, error) {
	result := r.db.Create(account)	
	if result.Error != nil {
		return nil, result.Error
	} 
	return account, nil
}

func (r *account) FindAccountByEmail(email string) (*model.Account, error) {
	var account model.Account
	result := r.db.First(&account, "email = ?", email)	
	if result.Error != nil {
		return nil, result.Error
	}	
	return &account, nil
}
