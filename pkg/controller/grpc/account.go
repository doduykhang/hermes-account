package grpc

import (
	"context"
	"doduykhang/hermes-account/internal/proto"
	"doduykhang/hermes-account/pkg/dto"
	"doduykhang/hermes-account/pkg/service"
)

type AccountServer struct {
	proto.UnimplementedAccountServiceServer
	service service.Account
}

func NewAccountServer(service service.Account) *AccountServer {
	return &AccountServer{
		service: service,	
	}
}

func (a *AccountServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	userInfo := req.GetUserInfo()

	var request dto.RegisterRequest
	request.Email = req.Email
	request.Password = req.Password
	request.UserInfo = dto.UserInfo{
		FirstName: userInfo.FirstName,
		LastName: userInfo.LastName,
	}
	
	userID, err := a.service.Register(request)

	if err != nil {
		return nil, err
	}

	var res proto.RegisterResponse
	res.UserID = userID
	return &res, nil
}

func (a *AccountServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var request dto.LoginRequest
	request.Email = req.Email
	request.Password = req.Password

	userID, err := a.service.Login(request)

	if err != nil {
		return nil, err
	}
	var res proto.LoginResponse
	res.UserID = userID
	return &res, nil
}
