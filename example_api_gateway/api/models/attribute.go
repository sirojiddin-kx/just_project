package models

type Attribute struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Attribute_Types string `json:"attribute_types"`
}

type CreateAttribute struct {
	Name            string `json:"name"`
	Attribute_Types string `json:"attribute_types"`
}

type AttributeId struct {
	Id string `json:"attribute_types"`
}

type GetAllAttributeResponse struct {
	Attributes []Attribute `json:"attributes"`
	Count      uint32      `json:"count"`
}

type UpdateAttribute struct {
	Resp string `json:"resp"`
}
