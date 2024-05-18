package controller

import (
	"errors"
	"net/http"
)

// GetAccessRights
// @Summary Get access rights
// @Security ApiKeyAuth
// @Tags AccessRight
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/diagrams/{diagram_id}/access [get]
func (c *Controller) GetAccessRights(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// GiveAccessRights
// @Summary Give access rights
// @Security ApiKeyAuth
// @Tags AccessRight
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/diagrams/{diagram_id}/access [post]
func (c *Controller) GiveAccessRights(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// DeleteAccessRights
// @Summary Delete access rights
// @Security ApiKeyAuth
// @Tags AccessRight
// @Accept json
// @Produce json
// @Param access_id path int true "Access ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/diagrams/access/{access_id} [delete]
func (c *Controller) DeleteAccessRights(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
