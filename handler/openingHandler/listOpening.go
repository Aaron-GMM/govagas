package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/Aaron-GMM/govagas/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openins := &[]schemas.Opening{}
	if err := handler.Db.Find(openins).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*openins) == 0 {
		handler.SendError(ctx, http.StatusNotFound, "[]")
		return
	}
	handler.SendSuccess(ctx, "list-openings", openins)
}
