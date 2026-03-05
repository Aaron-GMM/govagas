package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/gin-gonic/gin"
)

func (h *OpeningHandler) ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening, err := h.service.GetOpening(id)
	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	handler.SendSuccess(ctx, "show-opening", opening)
}
