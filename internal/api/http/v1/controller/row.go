package controller

import (
	"errors"
	"net/http"
)

// CreateRow
// @Summary Create row
// @Security ApiKeyAuth
// @Tags Row
// @Accept json
// @Produce json
// @Param table_id path int true "Table ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 201 {object} entity.NotImplemented
// @Router /v1/diagrams/tables/{table_id}/rows [post]
func (c *Controller) CreateRow(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// EditRow
// @Summary Edit row
// @Security ApiKeyAuth
// @Tags Row
// @Accept json
// @Produce json
// @Param row_id path int true "Row ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/tables/rows/{row_id} [put]
func (c *Controller) EditRow(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// DeleteRow
// @Summary Delete row
// @Security ApiKeyAuth
// @Tags Row
// @Accept json
// @Produce json
// @Param row_id path int true "Row ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/tables/rows/{row_id} [delete]
func (c *Controller) DeleteRow(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
