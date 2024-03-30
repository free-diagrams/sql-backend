package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetColors
// @Summary Get all colors
// @Security ApiKeyAuth
// @Tags Color
// @Description Get all colors
// @Accept json
// @Produce json
// @Param Accept-Language header string false "Language preference" default(en-US)
// @Success 200 {array} usecase.GetColorsResponseItem
// @Failure 401
// @Failure 500 {object} httpresp.Error
// @Router /colors [get]
func (c *Controller) GetColors(ctx *gin.Context) {
	c.log.Debug().Msg("Getting colors")

	response, err := c.useCase.GetColors(ctx)
	if err != nil {
		c.log.Error().Err(err).Msg("Failed to get colors")
		httpCode, messageID := c.ErrorToHTTPStatusAndMessageID(err)
		ctx.JSON(httpCode, c.ConstructError(ctx, messageID, err))
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
