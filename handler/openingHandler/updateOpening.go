package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/Aaron-GMM/govagas/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := &handler.UpadeteOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.ErrorF("Validate request error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		handler.Logger.ErrorF("Get opening id error: %v", id)
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id",
			"queryParameter").Error())
		return
	}

	opening := &schemas.Opening{}

	if err := handler.Db.First(&opening, id).Error; err != nil {
		handler.Logger.ErrorF("Get opening error: %v", err.Error())
		handler.SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	if err := handler.Db.Model(&opening).Updates(*request).Error; err != nil {
		handler.Logger.ErrorF("Update opening error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, "Update opening error")
		return
	}

	handler.SendSuccess(ctx, "update-opening", *opening)

}
