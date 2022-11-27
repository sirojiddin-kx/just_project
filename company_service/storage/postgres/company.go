package postgres

import (
	"bitbucket.org/udevs/example_api_gateway/genproto/company_service"
	"bitbucket.org/udevs/example_api_gateway/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type companyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) repo.CompanyRepoI {
	return &companyRepo{
		db: db,
	}
}

func (r *companyRepo) Create(req *company_service.CreateCompany) (string, error) {
	var (
		id uuid.UUID
	)
	tx, err := r.db.Begin()

	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	id, err = uuid.NewRandom()
	if err != nil {
		return "", err
	}

	query := `INSERT INTO company (id, name) VALUES($1, $2)`

	_, err = tx.Exec(query, id, req.Name)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *companyRepo) Get(id string) (*company_service.Company, error) {
	var company company_service.Company

	query := `SELECT id, name FROM company WHERE id = $1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&company.Id,
		&company.Name,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *companyRepo) GetAll(req *company_service.GetAllCompanyRequest) (*company_service.GetAllCompanyResponse, error) {
	var (
		args      = make(map[string]interface{})
		filter    string
		companies []*company_service.Company
		count     uint32
	)

	if req.Name != "" {
		filter += ` AND name ilike '%' || :name || '%' `
		args["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM company WHERE true ` + filter
	rows, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&count,
		)

		if err != nil {
			return nil, err
		}
	}

	filter += " OFFSET :offset LIMIT :limit "
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT id, name FROM company WHERE true ` + filter

	rows, err = r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var company company_service.Company

		err = rows.Scan(&company.Id, &company.Name)

		if err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	return &company_service.GetAllCompanyResponse{
		Companies: companies,
		Count:     count,
	}, nil
}

func (r *companyRepo) Update(req *company_service.Company) (string, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `UPDATE company SET name = $1 WHERE id = $2`
	_, err = tx.Exec(query, req.Name, req.Id)
	if err != nil {
		return "", err
	}

	return "updated", nil
}

func (r *companyRepo) Delete(req *company_service.CompanyId) (string, error) {
	query := `DELETE FROM company WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}
