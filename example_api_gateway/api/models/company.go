package models

type Company struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateCompany struct {
	Name string `json:"name"`
}

type CompanyId struct {
	Id string `json:"id"`
}

type GetAllCompanyResponse struct {
	Companies []Company `json:"companies"`
	Count     uint32    `json:"count"`
}

type UpdateCompany struct {
	Resp string `json:"resp"`
}
