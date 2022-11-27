package postgres

import (
	"bitbucket.org/Udevs/position_service/genproto/position_service"
	"bitbucket.org/Udevs/position_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type professionRepo struct {
	db *sqlx.DB
}

func NewProfessionRepo(db *sqlx.DB) repo.ProfessionRepoI {
	return &professionRepo{
		db: db,
	}
}

func (r *professionRepo) Create(req *position_service.CreateProfession) (string, error) {
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

	query := `
		INSERT INTO 
			profession
			(
				id,
				name
			)
			VALUES($1, $2)
	`

	_, err = tx.Exec(query, id, req.Name)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *professionRepo) Get(id string) (*position_service.Profession, error) {
	var profession position_service.Profession

	query := `
		SELECT 
			id,
			name
		FROM 
			profession
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&profession.Id,
		&profession.Name,
	)

	if err != nil {
		return nil, err
	}

	return &profession, nil
}

func (r *professionRepo) GetAll(req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error) {
	var (
		args        = make(map[string]interface{})
		filter      string
		professions []*position_service.Profession
		count       uint32
	)

	if req.Name != "" {
		filter += ` AND name ilike '%' || :name || '%' `
		args["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM profession WHERE true ` + filter
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

	query := `
		SELECT 
			id, 
			name
		FROM 
			profession WHERE true ` + filter

	rows, err = r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var profesion position_service.Profession

		err = rows.Scan(
			&profesion.Id,
			&profesion.Name,
		)

		if err != nil {
			return nil, err
		}

		professions = append(professions, &profesion)
	}

	return &position_service.GetAllProfessionResponse{
		Professions: professions,
		Count:       count,
	}, nil
}

func (r *professionRepo) Update(req *position_service.Profession) (string, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return "", err
	}

	// fmt.Printf("\n\n%+v\n\n", req)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `UPDATE profession SET name = $1 WHERE id = $2`
	_, err = tx.Exec(query, req.Name, req.Id)
	if err != nil {
		return "", err
	}

	return "updated", nil
}

func (r *professionRepo) Delete(req *position_service.ProfessionId) (string, error) {
	query := `DELETE FROM profession WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}
