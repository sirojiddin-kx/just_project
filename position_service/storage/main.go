package storage

import (
	"bitbucket.org/Udevs/position_service/storage/postgres"
	"bitbucket.org/Udevs/position_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Profession() repo.ProfessionRepoI
	Attribute() repo.AttributeRepoI
	Position() repo.PositionRepoI
}

type storagePG struct {
	profession repo.ProfessionRepoI
	attribute  repo.AttributeRepoI
	position   repo.PositionRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		profession: postgres.NewProfessionRepo(db),
		attribute:  postgres.NewAttributeRepo(db),
		position:   postgres.NewPositionRepo(db),
	}
}

func (s *storagePG) Profession() repo.ProfessionRepoI {
	return s.profession
}

func (s *storagePG) Attribute() repo.AttributeRepoI {
	return s.attribute
}

func (s *storagePG) Position() repo.PositionRepoI {
	return s.position
}
