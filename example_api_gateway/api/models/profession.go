package models

type CreateProfession struct {
	Name string `json:"name" binding:"required"`
}

type Profession struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetAllProfessionResponse struct {
	Professions []Profession `json:"professions"`
	Count       uint32       `json:"count"`
}

type UpdateProfession struct {
	Resp string `json:"resp"`
}
