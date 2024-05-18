package controller

import (
	"errors"
	"net/http"
)

// GetDatatypes
// @Summary Get datatypes
// @Security ApiKeyAuth
// @Tags Datatype
// @Accept json
// @Produce json
// @Param database_id path int true "Database ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/databases/{database_id}/datatypes [get]
func (c *Controller) GetDatatypes(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
