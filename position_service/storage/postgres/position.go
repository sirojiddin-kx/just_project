package postgres

import (
	"bitbucket.org/Udevs/position_service/genproto/position_service"
	"bitbucket.org/Udevs/position_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type positionRepo struct {
	db *sqlx.DB
}

func NewPositionRepo(db *sqlx.DB) repo.PositionRepoI {
	return &positionRepo{
		db: db,
	}
}

func (r *positionRepo) Create(req *position_service.CreatePositionRequest) (string, error) {
	var (
		idP uuid.UUID
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

	idP, err = uuid.NewRandom()
	if err != nil {
		return "", err
	}

	queryP := `INSERT INTO position (id, name, profession_id, company_id) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(queryP, idP, req.Name, req.ProfessionId, req.CompanyId)
	if err != nil {
		return "", err
	}

	queryPA := `INSERT INTO position_attribute (id, attribute_id, position_id, value) VALUES ($1, $2, $3, $4)`
	for _, val := range req.PositionAttributes {
		var idPA uuid.UUID
		idPA, err = uuid.NewRandom()
		if err != nil {
			return "", err
		}
		_, err = tx.Exec(queryPA, idPA, val.AttributeId, idP, val.Value)
		if err != nil {
			return "", err
		}
	}

	return idP.String(), nil
}

func (r *positionRepo) Get(id string) (*position_service.Position, error) {
	var position position_service.Position
	var position_attribute position_service.GetPositionAttribute
	var attribute position_service.Attribute

	queryP := `SELECT id, name, profession_id, company_id FROM position WHERE id = $1`

	row := r.db.QueryRow(queryP, id)
	err := row.Scan(
		&position.Id,
		&position.Name,
		&position.ProfessionId,
		&position.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	queryPA := `SELECT id, attribute_id, position_id, value FROM position_attribute
	WHERE position_id = $1`

	rows, err := r.db.Query(queryPA, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&position_attribute.Id,
			&position_attribute.AttributeId,
			&position_attribute.PositionId,
			&position_attribute.Value,
		)

		if err != nil {
			return nil, err
		}

		queryAtt := `SELECT id, name, current_state FROM attribute WHERE id = $1`

		row := r.db.QueryRow(queryAtt, position_attribute.AttributeId)
		err = row.Scan(
			&attribute.Id,
			&attribute.Name,
			&attribute.AttributeTypes,
		)

		if err != nil {
			return nil, err
		}

		position_attribute.Attribute = &attribute
		position.PositionAttributes = append(position.PositionAttributes, &position_attribute)

		defer rows.Close()
	}

	return &position, nil
}

func (r *positionRepo) GetAll(req *position_service.GetAllPositionRequest) (*position_service.GetAllPositionResponse, error) {
	var (
		args      = make(map[string]interface{})
		filter    string
		positions []*position_service.Position
		count     uint32
	)

	if req.Name != "" {
		filter += ` AND name ilike '%' || :name || '%' `
		args["name"] = req.Name
	}
	if req.ProfessionId != "" {
		filter += ` AND profession_id = :profession_id `
		args["profession_id"] = req.ProfessionId
	}
	if req.CompanyId != "" {
		filter += ` AND company_id = :company_id `
		args["company_id"] = req.CompanyId
	}
	countQuery := `SELECT count(1) FROM position WHERE true ` + filter
	row, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(
			&count,
		)

		if err != nil {
			return nil, err
		}
	}

	filter += " OFFSET :offset LIMIT :limit "
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT id, name, profession_id, company_id FROM position WHERE true ` + filter
	rows, err := r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var position position_service.Position
		err = rows.Scan(
			&position.Id,
			&position.Name,
			&position.ProfessionId,
			&position.CompanyId,
		)
		if err != nil {
			return nil, err
		}
		queryPA := `SELECT id, attribute_id, position_id, value FROM position_attribute
			WHERE position_id = $1`

		resp, err := r.db.Query(queryPA, position.Id)
		if err != nil {
			return nil, err
		}
		for resp.Next() {
			var position_attribute position_service.GetPositionAttribute
			err = resp.Scan(
				&position_attribute.Id,
				&position_attribute.AttributeId,
				&position_attribute.PositionId,
				&position_attribute.Value,
			)

			if err != nil {
				return nil, err
			}

			queryAtt := `SELECT id, name, current_state FROM attribute WHERE id = $1`

			res := r.db.QueryRow(queryAtt, position_attribute.AttributeId)
			var attribute position_service.Attribute
			err = res.Scan(
				&attribute.Id,
				&attribute.Name,
				&attribute.AttributeTypes,
			)

			if err != nil {
				return nil, err
			}
			position_attribute.Attribute = &attribute
			position.PositionAttributes = append(position.PositionAttributes, &position_attribute)
		}
		positions = append(positions, &position)
	}

	return &position_service.GetAllPositionResponse{
		Positions: positions,
		Count:     count,
	}, nil
}

func (r *positionRepo) Update(req *position_service.UpdatePositionRequest) (string, error) {
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

	query1 := `UPDATE position SET name = $1, profession_id = $2, company_id = $3 WHERE id = $4`
	_, err1 := tx.Exec(query1, req.Name, req.ProfessionId, req.CompanyId, req.Id)
	if err1 != nil {
		return "not updated", err1
	}

	query2 := `UPDATE position_attribute SET value = $1 WHERE id = $2`
	for _, val := range req.PositionAttribute {
		_, err = tx.Exec(query2, val.Value, val.Id)
		if err != nil {
			return "not updated", err
		}
	}
	query3 := `UPDATE attribute SET name = $1, current_state = $2 WHERE id = $3`
	for _, val := range req.Attribute {
		_, err = tx.Exec(query3, val.Name, val.AttributeTypes, val.Id)
		if err != nil {
			return "not updated", err
		}
	}

	return "updated", nil
}

func (r *positionRepo) Delete(req *position_service.PositionId) (string, error) {
	query := `DELETE FROM position_attribute WHERE position_id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return "", err
	}
	query = `DELETE FROM position WHERE id = $1`
	_, err = r.db.Exec(query, req.Id)
	if err != nil {
		return "", err
	}
	return "deleted", nil
}
