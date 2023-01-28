package service

import (
	"context"
	"doduykhang/hermes-account/pkg/dto"
	"doduykhang/hermes-account/pkg/model"
	"doduykhang/hermes-account/pkg/repository"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
)

type Account interface {
	Register(request dto.RegisterRequest) (string, error)
	Login(request dto.LoginRequest) (string, error)
}

type account struct {
	repo repository.Account
	rabbitMQ *amqp.Connection
}

func NewAccount(repo repository.Account, rabbitMQ *amqp.Connection) Account {
	return &account{
		repo: repo,
		rabbitMQ: rabbitMQ,
	}
}

func (a *account) publisCreateUserEvent(request dto.RegisterRequest, userID string) (error) {
	ch, err := a.rabbitMQ.Channel()
	defer ch.Close()

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
  		"create-user", // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)

	if err != nil {
		return err
	}

	var req struct {
		ID string `json:"id"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Email string `json:"email"`
	}

	req.ID = userID
	req.FirstName = request.UserInfo.FirstName
	req.LastName = request.UserInfo.LastName
	req.Email = request.Email

	body, err := json.Marshal(&req)

	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(),
  		"",     // exchange
  		q.Name, // routing key
  		false,  // mandatory
  		false,  // immediate
  		amqp.Publishing {
    			ContentType: "text/plain",
    			Body:        []byte(body),
  	})

	return err
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
	a.publisCreateUserEvent(request, account.UserID)
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


