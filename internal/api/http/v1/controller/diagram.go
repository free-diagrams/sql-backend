package controller

import (
	"errors"
	"net/http"
)

// GetRequestersDiagrams
// @Summary Get requesters diagrams
// @Security ApiKeyAuth
// @Tags Diagram
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/diagrams [get]
func (c *Controller) GetRequestersDiagrams(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// CreateDiagram
// @Summary Create diagram
// @Security ApiKeyAuth
// @Tags Diagram
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 201 {object} entity.NotImplemented
// @Router /v1/diagrams [post]
func (c *Controller) CreateDiagram(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// GetDiagram
// @Summary Get diagram
// @Security ApiKeyAuth
// @Tags Diagram
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/diagrams/{diagram_id} [get]
func (c *Controller) GetDiagram(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// EditDiagram
// @Summary Edit diagram
// @Security ApiKeyAuth
// @Tags Diagram
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/{diagram_id} [put]
func (c *Controller) EditDiagram(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// DeleteDiagram
// @Summary Delete diagram
// @Security ApiKeyAuth
// @Tags Diagram
// @Accept json
// @Produce json
// @Param diagram_id path int true "Diagram ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/{diagram_id} [delete]
func (c *Controller) DeleteDiagram(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
