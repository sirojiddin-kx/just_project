package service

import (
	"context"

	"bitbucket.org/udevs/example_api_gateway/genproto/company_service"
	"bitbucket.org/udevs/example_api_gateway/pkg/helper"
	"bitbucket.org/udevs/example_api_gateway/pkg/logger"
	"bitbucket.org/udevs/example_api_gateway/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
)

type companyService struct {
	logger  logger.Logger
	storage storage.StorageI
	company_service.UnimplementedCompanyServiceServer
}

func NewCompanyService(log logger.Logger, db *sqlx.DB) *companyService {
	return &companyService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (s *companyService) Create(ctx context.Context, req *company_service.CreateCompany) (*company_service.CompanyId, error) {
	id, err := s.storage.Company().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating company ", req, codes.Internal)
	}

	return &company_service.CompanyId{
		Id: id,
	}, nil
}

func (s *companyService) Get(ctx context.Context, req *company_service.CompanyId) (*company_service.Company, error) {
	company, err := s.storage.Company().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting company ", req, codes.Internal)
	}

	return company, nil
}

func (s *companyService) GetAll(ctx context.Context, req *company_service.GetAllCompanyRequest) (*company_service.GetAllCompanyResponse, error) {
	companies, err := s.storage.Company().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all company ", req, codes.Internal)
	}

	return companies, nil
}

func (s *companyService) Update(ctx context.Context, req *company_service.Company) (*company_service.UpdateCompany, error) {
	company, err := s.storage.Company().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting company ", req, codes.Internal)
	}
	return &company_service.UpdateCompany{
		Resp: company,
	}, nil
}

func (s *companyService) Delete(ctx context.Context, req *company_service.CompanyId) (*company_service.UpdateCompany, error) {
	company, err := s.storage.Company().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting company ", req, codes.Internal)
	}
	return &company_service.UpdateCompany{
		Resp: company,
	}, nil
}
