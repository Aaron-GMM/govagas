package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/Aaron-GMM/govagas/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "string").Error())
		return
	}
	opening := &schemas.Opening{}
	err := handler.Db.First(&opening, id).Error
	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if err := handler.Db.Delete(opening).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccess(ctx, "delete-opening", *opening)
}
