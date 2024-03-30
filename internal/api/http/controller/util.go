package controller

import (
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/free-diagrams/sql-backend/pkg/httpresp"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (c *Controller) ConstructError(ctx *gin.Context, messageID string, err error) *httpresp.Error {
	lang := ctx.MustGet("lang").(string)

	message, errLoc := c.localizer.TryLocalize(lang, messageID)
	if errLoc != nil {
		c.log.Debug().Err(errLoc).Msg("failed to localize message")
		message = c.localizer.English(messageID)
	}

	return &httpresp.Error{
		Message: message,
		Error:   err.Error(),
	}
}

func (c *Controller) ErrorToHTTPStatusAndMessageID(err error) (int, string) {

	if errors.Is(err, errs.InternalDatabase) {
		return http.StatusInternalServerError, "InternalDatabaseError"
	} else {
		return http.StatusInternalServerError, "UnspecifiedError"
	}
}
