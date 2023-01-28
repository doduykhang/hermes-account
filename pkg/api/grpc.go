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
	conf := config.LoadConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		log.Fatal(err)
	}
	
	s := grpc.NewServer()

	//config
	db := config.NewGorm(conf)
	db.AutoMigrate(&model.Account{})

	rabbitMq := config.NewRabbitMq(conf)
	defer rabbitMq.Close()
	//repo
	accountRepo := repository.NewAccount(db)

	//service
	accountService := service.NewAccount(accountRepo, rabbitMq)

	//server 
	accountServer := controller.NewAccountServer(accountService)

	//register server
	proto.RegisterAccountServiceServer(s, accountServer)

	log.Printf("grpc server starting on port %s", conf.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
