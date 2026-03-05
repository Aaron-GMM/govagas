package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/gin-gonic/gin"
)

func (h *OpeningHandler) DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "string").Error())
		return
	}

	err := h.service.DeleteOpening(id)
	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	handler.SendSuccess(ctx, "delete-opening", http.StatusNoContent)
}
