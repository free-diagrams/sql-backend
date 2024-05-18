package controller

import (
	"errors"
	"net/http"
)

// Login
// @Summary Login
// @Tags Authorization
// @Accept json
// @Produce json
// @Param req body entity.NotImplemented true "Request body"
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/auth/login [post]
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// Refresh
// @Summary Refresh
// @Tags Authorization
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/auth/refresh [post]
func (c *Controller) Refresh(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}

// Logout
// @Summary Logout
// @Tags Authorization
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {object} entity.NotImplemented
// @Router /v1/auth/logout [post]
func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Не реализовано")
}
