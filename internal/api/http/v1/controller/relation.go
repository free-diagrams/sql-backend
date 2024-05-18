package controller

import (
	"errors"
	"net/http"
)

// CreateRelation
// @Summary Create relation
// @Security ApiKeyAuth
// @Tags Relation
// @Accept json
// @Produce json
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 201 {object} entity.NotImplemented
// @Router /v1/diagrams/tables/relations [post]
func (c *Controller) CreateRelation(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// EditRelation
// @Summary Edit relation
// @Security ApiKeyAuth
// @Tags Relation
// @Accept json
// @Produce json
// @Param relation_id path int true "Relation ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/tables/relations/{relation_id} [put]
func (c *Controller) EditRelation(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// DeleteRelation
// @Summary Delete relation
// @Security ApiKeyAuth
// @Tags Relation
// @Accept json
// @Produce json
// @Param relation_id path int true "Relation ID"
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/diagrams/tables/relations/{relation_id} [delete]
func (c *Controller) DeleteRelation(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
