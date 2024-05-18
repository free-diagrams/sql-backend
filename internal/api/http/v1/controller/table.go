package controller

import (
	"errors"
	"net/http"
)

// CreateTable
// @Summary Create table
// @Security ApiKeyAuth
// @Tags Table
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 201 {object} entity.NotImplemented
// @Router /v1/diagrams/{diagram_id}/tables [post]
func (c *Controller) CreateTable(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// EditTable
// @Summary Edit table
// @Security ApiKeyAuth
// @Tags Table
// @Accept json
// @Produce json
// @Param table_id path int true "Table ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/tables/{table_id} [put]
func (c *Controller) EditTable(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// DeleteTable
// @Summary Delete table
// @Security ApiKeyAuth
// @Tags Table
// @Accept json
// @Produce json
// @Param table_id path int true "Table ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/tables/{table_id} [delete]
func (c *Controller) DeleteTable(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
