package controller

import (
	"errors"
	"net/http"
)

// GetRequesterUser
// @Summary Get requester user
// @Security ApiKeyAuth
// @Tags User
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/users/me [get]
func (c *Controller) GetRequesterUser(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// EditRequesterUser
// @Summary Edit requester user
// @Security ApiKeyAuth
// @Tags User
// @Accept json
// @Produce json
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 204
// @Router /v1/users/me [put]
func (c *Controller) EditRequesterUser(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// GetUser
// @Summary Get user
// @Security ApiKeyAuth
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/users/{user_id} [get]
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
