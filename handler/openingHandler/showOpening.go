package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/Aaron-GMM/govagas/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := &schemas.Opening{}
	if err := handler.Db.First(&opening, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	handler.SendSuccess(ctx, "show-opening", *opening)
}
