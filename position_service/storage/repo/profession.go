package repo

import "bitbucket.org/Udevs/position_service/genproto/position_service"

type ProfessionRepoI interface {
	Create(req *position_service.CreateProfession) (string, error)
	Get(id string) (*position_service.Profession, error)
	GetAll(req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error)
	Update(req *position_service.Profession) (string, error)
	Delete(req *position_service.ProfessionId) (string, error)
}
