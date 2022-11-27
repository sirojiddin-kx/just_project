package services

import (
	"fmt"

	"bitbucket.org/udevs/example_api_gateway/config"
	"bitbucket.org/udevs/example_api_gateway/genproto/company_service"
	"bitbucket.org/udevs/example_api_gateway/genproto/position_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	ProfessionService() position_service.ProfessionServiceClient
	AttributeService() position_service.AttributeServiceClient
	PositionService() position_service.PositionServiceClient
	CompanyService() company_service.CompanyServiceClient
}

type grpcClients struct {
	professionService position_service.ProfessionServiceClient
	attributeService  position_service.AttributeServiceClient
	positionService   position_service.PositionServiceClient
	companyService    company_service.CompanyServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connPositionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connCompanyService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CompanyServiceHost, conf.CompanyServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		professionService: position_service.NewProfessionServiceClient(connPositionService),
		attributeService:  position_service.NewAttributeServiceClient(connPositionService),
		positionService:   position_service.NewPositionServiceClient(connPositionService),
		companyService:    company_service.NewCompanyServiceClient(connCompanyService),
	}, nil
}

func (g *grpcClients) ProfessionService() position_service.ProfessionServiceClient {
	return g.professionService
}

func (g *grpcClients) AttributeService() position_service.AttributeServiceClient {
	return g.attributeService
}

func (g *grpcClients) PositionService() position_service.PositionServiceClient {
	return g.positionService
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}
