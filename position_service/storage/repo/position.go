package repo

import "bitbucket.org/Udevs/position_service/genproto/position_service"

type PositionRepoI interface {
	Create(req *position_service.CreatePositionRequest) (string, error)
	Get(id string) (*position_service.Position, error)
	GetAll(req *position_service.GetAllPositionRequest) (*position_service.GetAllPositionResponse, error)
	Update(req *position_service.UpdatePositionRequest) (string, error)
	Delete(req *position_service.PositionId) (string, error)
}
