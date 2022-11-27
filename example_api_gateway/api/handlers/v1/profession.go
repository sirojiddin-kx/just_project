package v1

import (
	"context"
	"errors"
	"net/http"

	"bitbucket.org/udevs/example_api_gateway/api/models"
	"bitbucket.org/udevs/example_api_gateway/genproto/position_service"
	"bitbucket.org/udevs/example_api_gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// Create Profession godoc
// @ID create-profession
// @Router /v1/profession [POST]
// @Summary create profession
// @Description Create Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body models.CreateProfession true "profession"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfession(c *gin.Context) {
	var profession models.CreateProfession

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.ProfessionService().Create(
		context.Background(),
		&position_service.CreateProfession{
			Name: profession.Name,
		},
	)

	if !handleError(h.log, c, err, "error while creating profession") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Profession godoc
// @ID get-profession
// @Router /v1/profession/{profession_id} [GET]
// @Summary get profession
// @Description Get Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession_id path string true "profession_id"
// @Success 200 {object} models.ResponseModel{data=models.Profession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetProfession(c *gin.Context) {
	var profession models.Profession
	profession_id := c.Param("profession_id")

	if !util.IsValidUUID(profession_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "profession id is not valid", errors.New("profession id is not valid"))
		return
	}

	resp, err := h.services.ProfessionService().Get(
		context.Background(),
		&position_service.ProfessionId{
			Id: profession_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting profession") {
		return
	}

	err = ParseToStruct(&profession, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", profession)
}

// Get All Profession godoc
// @ID get-all-profession
// @Router /v1/profession [GET]
// @Summary get all profession
// @Description Get All Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProfessions(c *gin.Context) {
	var professions models.GetAllProfessionResponse

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.ProfessionService().GetAll(
		context.Background(),
		&position_service.GetAllProfessionRequest{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Name:   c.Query("name"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all professions") {
		return
	}

	err = ParseToStruct(&professions, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", professions)
}

// Update Profession godoc
// @ID update-profession
// @Router /v1/profession/update/ [PUT]
// @Summary update profession
// @Description Update Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body models.Profession true "profession"
// @Success 200 {object} models.ResponseModel{data=models.UpdateProfession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateProfession(c *gin.Context) {
	var profession models.Profession

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.ProfessionService().Update(
		context.Background(),
		&position_service.Profession{
			Name: profession.Name,
			Id:   profession.ID,
		},
	)

	if !handleError(h.log, c, err, "error while updating profession") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Profession godoc
// @ID delete-profession
// @Router /v1/profession/delete/{profession_id} [DELETE]
// @Summary delete profession
// @Description Delete Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession_id path string true "profession_id"
// @Success 200 {object} models.ResponseModel{data=models.UpdateProfession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteProfession(c *gin.Context) {
	profession_id := c.Param("profession_id")

	if !util.IsValidUUID(profession_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "profession id is not valid", errors.New("profession id is not valid"))
		return
	}

	resp, err := h.services.ProfessionService().Delete(
		context.Background(),
		&position_service.ProfessionId{
			Id: profession_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting profession") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}
