package repo

import "bitbucket.org/Udevs/position_service/genproto/position_service"

type AttributeRepoI interface {
	Create(req *position_service.CreateAttribute) (string, error)
	Get(id string) (*position_service.Attribute, error)
	GetAll(req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error)
	Update(req *position_service.Attribute) (string, error)
	Delete(req *position_service.AttributeId) (string, error)
}
