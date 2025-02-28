package order

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/mehgokalp/re-partners-challenge/internal/packaging"
	"github.com/rotisserie/eris"
	"net/http"
)

type Handler struct {
	ctx              *gin.Context
	logger           log.Logger
	packagingHandler *packaging.Handler
}

// NewHandler godoc
// TODO REGENERATE THIS
// @Summary Update order status
// @Description Calculate order packaging
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {object} nil
// @Failure 406 {object} error
// @Failure 500 {object} error
// @Router /messages/ [patch]
func NewHandler(logger log.Logger, packagingHandler *packaging.Handler) func(*gin.Context) {
	return func(c *gin.Context) {
		h := Handler{
			ctx:              c,
			logger:           logger,
			packagingHandler: packagingHandler,
		}

		h.Handle()
	}
}

type orderPackagingForm struct {
	Items int `form:"items" binding:"gte=1"`
}

func (h *Handler) Handle() {
	var form orderPackagingForm
	err := h.ctx.ShouldBindQuery(&form)

	if err != nil {
		h.logger.Errorf(eris.Wrap(err, "failed to bind json").Error())

		h.ctx.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	packStack, err := h.packagingHandler.Pack(form.Items)
	if err != nil {
		h.logger.Errorf(eris.Wrap(err, "failed to pack").Error())

		h.ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	h.ctx.JSON(http.StatusOK, MapPacks(packStack))
}
