package controller

import (
	"net/http"
)

type GetDatabasesResponseItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetDatabases
// @Summary Get databases
// @Security ApiKeyAuth
// @Tags Database
// @Produce json
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {array} GetDatabasesResponseItem
// @Router /v1/databases [get]
func (c *Controller) GetDatabases(w http.ResponseWriter, r *http.Request) error {
	c.log.Debug().Msg("Getting all databases")

	databases, err := c.databaseUseCase.GetAllDatabases(r.Context())
	if err != nil {
		return err
	}
	databasesResponse := make([]GetDatabasesResponseItem, len(databases))
	for i, database := range databases {
		databasesResponse[i] = GetDatabasesResponseItem{
			ID:   *database.ID,
			Name: database.Name,
		}
	}

	c.responseWriter.Write(r.Context(), w, http.StatusOK, databasesResponse)
	return nil
}
