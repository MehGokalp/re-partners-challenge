package start_stop

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/rotisserie/eris"
	"net/http"
)

type Handler struct {
	ctx    *gin.Context
	logger log.Logger
}

// NewHandler godoc
// @Summary Update message status
// @Description Update the status of the message engine
// @Tags messages
// @Accept  json
// @Produce  json
// @Param action body startStopForm true "Action to start or stop the message engine"
// @Success 202 {object} nil
// @Failure 406 {object} error
// @Failure 500 {object} error
// @Router /messages/ [patch]
func NewHandler(logger log.Logger) func(*gin.Context) {
	return func(c *gin.Context) {
		h := Handler{
			ctx:    c,
			logger: logger,
		}

		h.Handle()
	}
}

type startStopForm struct {
	Items int `json:"items" binding:"gte=1"`
}

func (h *Handler) Handle() {
	var form startStopForm
	err := h.ctx.BindJSON(&form)

	if err != nil {
		h.logger.Errorf(eris.Wrap(err, "failed to bind json").Error())

		h.ctx.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	h.ctx.JSON(http.StatusOK, nil)
}
