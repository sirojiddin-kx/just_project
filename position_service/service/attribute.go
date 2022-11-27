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

type attributeService struct {
	logger  logger.Logger
	storage storage.StorageI
	position_service.UnimplementedAttributeServiceServer
}

func NewAttributeService(log logger.Logger, db *sqlx.DB) *attributeService {
	return &attributeService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (s *attributeService) Create(ctx context.Context, req *position_service.CreateAttribute) (*position_service.AttributeId, error) {
	id, err := s.storage.Attribute().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating attribute ", req, codes.Internal)
	}

	return &position_service.AttributeId{
		Id: id,
	}, nil
}

func (s *attributeService) Get(ctx context.Context, req *position_service.AttributeId) (*position_service.Attribute, error) {
	attribute, err := s.storage.Attribute().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting attribute ", req, codes.Internal)
	}

	return attribute, nil
}

func (s *attributeService) GetAll(ctx context.Context, req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error) {
	attributes, err := s.storage.Attribute().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all attribute ", req, codes.Internal)
	}

	return attributes, nil
}

func (s *attributeService) Update(ctx context.Context, req *position_service.Attribute) (*position_service.UpdateAttribute, error) {
	attribute, err := s.storage.Attribute().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating attribute ", req, codes.Internal)
	}
	return &position_service.UpdateAttribute{
		Resp: attribute,
	}, nil
}

func (s *attributeService) Delete(ctx context.Context, req *position_service.AttributeId) (*position_service.UpdateAttribute, error) {
	attribute, err := s.storage.Attribute().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting attribute ", req, codes.Internal)
	}
	return &position_service.UpdateAttribute{
		Resp: attribute,
	}, nil
}
