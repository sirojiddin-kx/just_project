package main

import (
	"fmt"
	"net"

	"bitbucket.org/Udevs/position_service/config"
	"bitbucket.org/Udevs/position_service/genproto/position_service"
	"bitbucket.org/Udevs/position_service/pkg/logger"
	"bitbucket.org/Udevs/position_service/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "position_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)

	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	professionService := service.NewProfessionService(log, db)
	attributeService := service.NewAttributeService(log, db)
	positionService := service.NewPositionService(log, db)

	s := grpc.NewServer()
	reflection.Register(s)

	position_service.RegisterProfessionServiceServer(s, professionService)
	position_service.RegisterAttributeServiceServer(s, attributeService)
	position_service.RegisterPositionServiceServer(s, positionService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
