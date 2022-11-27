package storage

import (
	"bitbucket.org/udevs/example_api_gateway/storage/postgres"
	"bitbucket.org/udevs/example_api_gateway/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Company() repo.CompanyRepoI
}

type storagePG struct {
	company repo.CompanyRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		company: postgres.NewCompanyRepo(db),
	}
}

func (s *storagePG) Company() repo.CompanyRepoI {
	return s.company
}
