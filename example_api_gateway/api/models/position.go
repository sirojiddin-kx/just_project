package models

type PositionAttribute struct {
	AttributeId string `json:"attribute_id"`
	Value       string `json:"value"`
}

type PositionAttribute2 struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type GetPositionAttribute struct {
	Id          string    `json:"id"`
	AttributeId string    `json:"attribute_id"`
	PositionId  string    `json:"position_id"`
	Value       string    `json:"value"`
	Attribute   Attribute `json:"attribute"`
}

type Position struct {
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	ProfessionId       string                 `json:"profession_id"`
	CompanyId          string                 `json:"company_id"`
	PositionAttributes []GetPositionAttribute `json:"position_attributes"`
}

type CreatePositionRequest struct {
	Name               string              `json:"name"`
	ProfessionId       string              `json:"profession_id"`
	CompanyId          string              `json:"company_id"`
	PositionAttributes []PositionAttribute `json:"position_attributes"`
}

type GetAllPositionRequest struct {
	Limit        uint32 `json:"limit"`
	Offset       uint32 `json:"offset"`
	Name         string `json:"name"`
	ProfessionId string `json:"profession_id"`
	CompanyId    string `json:"company_id"`
}

type GetAllPositionResponse struct {
	Positions []Position `json:"positions"`
	Count     uint32     `json:"count"`
}

type PositionId struct {
	Id string `json:"id"`
}

type UpdatePosition struct {
	Resp string `json:"resp"`
}

type UpdatePositionRequest struct {
	Id                string               `json:"id"`
	Name              string               `json:"name"`
	ProfessionId      string               `json:"profession_id"`
	CompanyId         string               `json:"company_id"`
	PositionAttribute []PositionAttribute2 `json:"position_attribute"`
	Attribute         []Attribute          `json:"attribute"`
}
