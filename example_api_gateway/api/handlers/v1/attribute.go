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

// Create Attribute godoc
// @ID create-attribute
// @Router /v1/attribute [POST]
// @Summary create attribute
// @Description Create Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body models.CreateAttribute true "attribute"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateAttribute(c *gin.Context) {
	var attribute models.CreateAttribute

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.AttributeService().Create(
		context.Background(),
		&position_service.CreateAttribute{
			Name:           attribute.Name,
			AttributeTypes: attribute.Attribute_Types,
		},
	)

	if !handleError(h.log, c, err, "error while creating attribute") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Attribute godoc
// @ID get-attribute
// @Router /v1/attribute/{attribute_id} [GET]
// @Summary get attribute
// @Description Get Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Success 200 {object} models.ResponseModel{data=models.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAttribute(c *gin.Context) {
	var attribute models.Attribute
	attribute_id := c.Param("attribute_id")

	if !util.IsValidUUID(attribute_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "attribute id is not valid", errors.New("attribute id is not valid"))
		return
	}

	resp, err := h.services.AttributeService().Get(
		context.Background(),
		&position_service.AttributeId{
			Id: attribute_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting attribute") {
		return
	}

	err = ParseToStruct(&attribute, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", attribute)
}

// Get All Attribute godoc
// @ID get-all-attribute
// @Router /v1/attribute [GET]
// @Summary get all attribute
// @Description Get All Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllAttributes(c *gin.Context) {
	var attributes models.GetAllAttributeResponse

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.AttributeService().GetAll(
		context.Background(),
		&position_service.GetAllAttributeRequest{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Name:   c.Query("name"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all attributes") {
		return
	}

	err = ParseToStruct(&attributes, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", attributes)
}

// Update Attribute godoc
// @ID update-attribute
// @Router /v1/attribute/update/ [PUT]
// @Summary update attribute
// @Description Update Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body models.Attribute true "attribute"
// @Success 200 {object} models.ResponseModel{data=models.UpdateAttribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateAttribute(c *gin.Context) {
	var attribute models.Attribute

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.AttributeService().Update(
		context.Background(),
		&position_service.Attribute{
			Name:           attribute.Name,
			Id:             attribute.Id,
			AttributeTypes: attribute.Attribute_Types,
		},
	)

	if !handleError(h.log, c, err, "error while updating attribute") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Attribute godoc
// @ID delete-attribute
// @Router /v1/attribute/delete/{attribute_id} [DELETE]
// @Summary delete attribute
// @Description Delete Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Success 200 {object} models.ResponseModel{data=models.UpdateAttribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteAttribute(c *gin.Context) {
	attribute_id := c.Param("attribute_id")

	if !util.IsValidUUID(attribute_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "attribute id is not valid", errors.New("attribute id is not valid"))
		return
	}

	resp, err := h.services.AttributeService().Delete(
		context.Background(),
		&position_service.AttributeId{
			Id: attribute_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting attribute") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}
