package controller

import (
	"errors"
	"net/http"
)

// GetTableCoordinates
// @Summary Get table coordinates
// @Security ApiKeyAuth
// @Tags Coordinate
// @Accept json
// @Produce json
// @Param table_id path int true "Table ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/tables/{table_id}/coordinates [get]
func (c *Controller) GetTableCoordinates(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// EditTableCoordinates
// @Summary Get table coordinates
// @Security ApiKeyAuth
// @Tags Coordinate
// @Accept json
// @Produce json
// @Param table_id path int true "Table ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/tables/{table_id}/coordinates [put]
func (c *Controller) EditTableCoordinates(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
