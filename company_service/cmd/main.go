package main

import (
	"fmt"
	"net"

	"bitbucket.org/udevs/example_api_gateway/config"
	"bitbucket.org/udevs/example_api_gateway/genproto/company_service"
	"bitbucket.org/udevs/example_api_gateway/pkg/logger"
	"bitbucket.org/udevs/example_api_gateway/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "company_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)
	fmt.Println("Postgres conStr : ", conStr)
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		fmt.Println("Err ::::::::;", err)
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	companyService := service.NewCompanyService(log, db)

	s := grpc.NewServer()
	reflection.Register(s)

	company_service.RegisterCompanyServiceServer(s, companyService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
