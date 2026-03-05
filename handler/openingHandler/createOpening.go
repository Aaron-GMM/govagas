package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/Aaron-GMM/govagas/schemas"
	"github.com/gin-gonic/gin"
)

func (h *OpeningHandler) CreateOpeningHandler(ctx *gin.Context) {
	request := handler.CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %s", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Name:            request.Name,
		PublicationData: request.PublicationData,
		Role:            request.Role,
		Company:         request.Company,
		Locate:          request.Locate,
		Remote:          *request.Remote,
		Link:            request.Link,
		Salary:          request.Salary,
		Description:     request.Description,
	}

	createdOpening, err := h.service.CreateOpening(&opening)
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(ctx, "create-opening", createdOpening)
}
