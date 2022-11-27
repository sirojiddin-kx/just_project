package postgres

import (
	"bitbucket.org/Udevs/position_service/genproto/position_service"
	"bitbucket.org/Udevs/position_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type attributeRepo struct {
	db *sqlx.DB
}

func NewAttributeRepo(db *sqlx.DB) repo.AttributeRepoI {
	return &attributeRepo{
		db: db,
	}
}

func (r *attributeRepo) Create(req *position_service.CreateAttribute) (string, error) {
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

	query := `INSERT INTO attribute (id, name, current_state) VALUES($1, $2, $3)`

	_, err = tx.Exec(query, id, req.Name, req.AttributeTypes)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *attributeRepo) Get(id string) (*position_service.Attribute, error) {
	var attribute position_service.Attribute

	query := `SELECT id, name, current_state FROM attribute WHERE id = $1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&attribute.Id,
		&attribute.Name,
		&attribute.AttributeTypes,
	)

	if err != nil {
		return nil, err
	}

	return &attribute, nil
}

func (r *attributeRepo) GetAll(req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error) {
	var (
		args       = make(map[string]interface{})
		filter     string
		attributes []*position_service.Attribute
		count      uint32
	)

	if req.Name != "" {
		filter += ` AND name ilike '%' || :name || '%' `
		args["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM attribute WHERE true ` + filter
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

	query := `SELECT id, name, current_state FROM attribute WHERE true ` + filter

	rows, err = r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var attribute position_service.Attribute

		err = rows.Scan(
			&attribute.Id,
			&attribute.Name,
			&attribute.AttributeTypes,
		)

		if err != nil {
			return nil, err
		}

		attributes = append(attributes, &attribute)
	}

	return &position_service.GetAllAttributeResponse{
		Attributes: attributes,
		Count:      count,
	}, nil
}

func (r *attributeRepo) Update(req *position_service.Attribute) (string, error) {
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

	query := `UPDATE attribute SET name = $1, current_state = $2 WHERE id = $3`
	_, err = tx.Exec(query, req.Name, req.AttributeTypes, req.Id)
	if err != nil {
		return "", err
	}

	return "updated", nil
}

func (r *attributeRepo) Delete(req *position_service.AttributeId) (string, error) {
	query := `DELETE FROM attribute WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}
