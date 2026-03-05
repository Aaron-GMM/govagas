package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/gin-gonic/gin"
)

func (h *OpeningHandler) ListOpeningHandler(ctx *gin.Context) {
	openings, err := h.service.ListOpenings()
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if len(openings) == 0 {
		handler.SendError(ctx, http.StatusNoContent, "[]")
		return
	}
	handler.SendSuccess(ctx, "list-openings", openings)
}
