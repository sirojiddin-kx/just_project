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

// Create Position godoc
// @ID create-position
// @Router /v1/position [POST]
// @Summary create position
// @Description Create Position
// @Tags position
// @Accept json
// @Produce json
// @Param position body models.CreatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreatePositionRequest(c *gin.Context) {
	var position models.CreatePositionRequest
	var pos_att []*position_service.PositionAttribute

	if err := c.BindJSON(&position); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	for _, val := range position.PositionAttributes {
		pos_att = append(pos_att, &position_service.PositionAttribute{
			AttributeId: val.AttributeId,
			Value:       val.Value,
		})
	}

	resp, err := h.services.PositionService().Create(
		context.Background(),
		&position_service.CreatePositionRequest{
			Name:               position.Name,
			ProfessionId:       position.ProfessionId,
			CompanyId:          position.CompanyId,
			PositionAttributes: pos_att,
		},
	)

	if !handleError(h.log, c, err, "error while creating position") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Position godoc
// @ID get-position
// @Router /v1/position/{position_id} [GET]
// @Summary get position
// @Description Get Position
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string true "position_id"
// @Success 200 {object} models.ResponseModel{data=models.Position} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetPosition(c *gin.Context) {
	var position models.Position
	position_id := c.Param("position_id")

	if !util.IsValidUUID(position_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "position id is not valid", errors.New("position id is not valid"))
		return
	}

	resp, err := h.services.PositionService().Get(
		context.Background(),
		&position_service.PositionId{
			Id: position_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting position") {
		return
	}

	err = ParseToStruct(&position, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", position)
}

// Get All Position godoc
// @ID get-all-position
// @Router /v1/position [GET]
// @Summary get all position
// @Description Get All Position
// @Tags position
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param profession_id query string false "profession_id"
// @Param company_id query string false "company_id"
// @Success 200 {object} models.ResponseModel{data=models.GetAllPositionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllPositions(c *gin.Context) {
	var positions models.GetAllPositionResponse

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.PositionService().GetAll(
		context.Background(),
		&position_service.GetAllPositionRequest{
			Limit:        uint32(limit),
			Offset:       uint32(offset),
			Name:         c.Query("name"),
			ProfessionId: c.Query("profession_id"),
			CompanyId:    c.Query("company_id"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all positions") {
		return
	}

	err = ParseToStruct(&positions, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", positions)
}

// Update Position godoc
// @ID update-position
// @Router /v1/position/update/ [PUT]
// @Summary update position
// @Description Update Position
// @Tags position
// @Accept json
// @Produce json
// @Param position body models.UpdatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=models.UpdatePosition} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdatePosition(c *gin.Context) {
	var position models.UpdatePositionRequest
	var pos_att []*position_service.PositionAttribute2
	var att []*position_service.Attribute

	if err := c.BindJSON(&position); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	for _, val := range position.PositionAttribute {
		pos_att = append(pos_att, &position_service.PositionAttribute2{
			Id:    val.Id,
			Value: val.Value,
		})
	}

	for _, val := range position.Attribute {
		att = append(att, &position_service.Attribute{
			Id:             val.Id,
			Name:           val.Name,
			AttributeTypes: val.Attribute_Types,
		})
	}

	_, err := h.services.PositionService().Update(
		context.Background(),
		&position_service.UpdatePositionRequest{
			Id:                position.Id,
			Name:              position.Name,
			ProfessionId:      position.ProfessionId,
			CompanyId:         position.CompanyId,
			PositionAttribute: pos_att,
			Attribute:         att,
		},
	)

	if !handleError(h.log, c, err, "error while updating position") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", &models.UpdatePosition{Resp: "updated"})
}

// Delete Position godoc
// @ID delete-position
// @Router /v1/position/delete/{position_id} [DELETE]
// @Summary delete position
// @Description Delete Position
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string true "position_id"
// @Success 200 {object} models.ResponseModel{data=models.UpdatePosition} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeletePosition(c *gin.Context) {
	position_id := c.Param("position_id")

	if !util.IsValidUUID(position_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "position id is not valid", errors.New("position id is not valid"))
		return
	}

	resp, err := h.services.PositionService().Delete(
		context.Background(),
		&position_service.PositionId{
			Id: position_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting position") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
