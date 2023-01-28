package api

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"doduykhang/hermes-account/pkg/config"
	"doduykhang/hermes-account/pkg/model"
	"doduykhang/hermes-account/pkg/repository"
	"doduykhang/hermes-account/pkg/service"
	"doduykhang/hermes-account/internal/proto"
	controller "doduykhang/hermes-account/pkg/controller/grpc"
)

func GRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8081"))
	if err != nil {
		log.Fatal(err)
	}
	
	s := grpc.NewServer()

	dsn := "sammy:password@tcp(127.0.0.1:3306)/hermes_account?charset=utf8mb4&parseTime=True&loc=Local"
	db := config.NewGorm(dsn)
	
	db.AutoMigrate(&model.Account{})
	//repo
	accountRepo := repository.NewAccount(db)

	//service
	accountService := service.NewAccount(accountRepo)

	//server 
	accountServer := controller.NewAccountServer(accountService)

	//register server
	proto.RegisterAccountServiceServer(s, accountServer)

	log.Printf("grpc server starting on port %s", "8081")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
