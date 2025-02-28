package order

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/mehgokalp/re-partners-challenge/internal/packaging/domain"
	"github.com/rotisserie/eris"
	"net/http"
)

type Handler struct {
	ctx              *gin.Context
	logger           log.Logger
	packagingHandler *domain.Handler
}

// NewHandler
// @Summary Calculate order packaging
// @Description Calculate the best packaging option for a given number of items
// @Tags packaging
// @Accept  json
// @Produce  json
// @Param items query int true "Number of items" minimum(1)
// @Success 200 {array} packaging.Pack "List of packs"
// @Failure 406 {object} error "Invalid input"
// @Failure 500 {object} error "Internal server error"
// @Router /v1/calculate-packaging [get]
func NewHandler(logger log.Logger, packagingHandler *domain.Handler) func(*gin.Context) {
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
