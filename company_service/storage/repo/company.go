package repo

import "bitbucket.org/udevs/example_api_gateway/genproto/company_service"

type CompanyRepoI interface {
	Create(req *company_service.CreateCompany) (string, error)
	Get(id string) (*company_service.Company, error)
	GetAll(req *company_service.GetAllCompanyRequest) (*company_service.GetAllCompanyResponse, error)
	Update(req *company_service.Company) (string, error)
	Delete(req *company_service.CompanyId) (string, error)
}
