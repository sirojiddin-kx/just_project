package service

import (
	"context"

	"bitbucket.org/Udevs/position_service/genproto/position_service"
	"bitbucket.org/Udevs/position_service/pkg/helper"
	"bitbucket.org/Udevs/position_service/pkg/logger"
	"bitbucket.org/Udevs/position_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
)

type positionService struct {
	logger  logger.Logger
	storage storage.StorageI
	position_service.UnimplementedPositionServiceServer
}

func NewPositionService(log logger.Logger, db *sqlx.DB) *positionService {
	return &positionService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (s *positionService) Create(ctx context.Context, req *position_service.CreatePositionRequest) (*position_service.PositionId, error) {
	id, err := s.storage.Position().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating position ", req, codes.Internal)
	}

	return &position_service.PositionId{
		Id: id,
	}, nil
}

func (s *positionService) Get(ctx context.Context, req *position_service.PositionId) (*position_service.Position, error) {
	position, err := s.storage.Position().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting attribute ", req, codes.Internal)
	}

	return position, nil
}

func (s *positionService) GetAll(ctx context.Context, req *position_service.GetAllPositionRequest) (*position_service.GetAllPositionResponse, error) {
	positions, err := s.storage.Position().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all attribute ", req, codes.Internal)
	}

	return positions, nil
}

func (s *positionService) Update(ctx context.Context, req *position_service.UpdatePositionRequest) (*position_service.UpdatePosition, error) {
	position, err := s.storage.Position().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating position ", req, codes.Internal)
	}
	return &position_service.UpdatePosition{
		Resp: position,
	}, nil
}

func (s *positionService) Delete(ctx context.Context, req *position_service.PositionId) (*position_service.UpdatePosition, error) {
	position, err := s.storage.Position().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting position ", req, codes.Internal)
	}
	return &position_service.UpdatePosition{
		Resp: position,
	}, nil
}
