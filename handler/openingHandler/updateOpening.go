package openingHandler

import (
	"net/http"

	"github.com/Aaron-GMM/govagas/handler"
	"github.com/gin-gonic/gin"
)

func (h *OpeningHandler) UpdateOpeningHandler(ctx *gin.Context) {
	request := &handler.UpadeteOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	// 1. Busca os dados atuais usando o Service
	opening, err := h.service.GetOpening(id)
	if err != nil {
		handler.SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	// 2. Atualiza apenas os campos fornecidos no request
	if request.Name != nil {
		opening.Name = *request.Name
	}
	if request.PublicationData != nil {
		opening.PublicationData = *request.PublicationData
	}
	if request.Role != nil {
		opening.Role = *request.Role
	}
	if request.Company != nil {
		opening.Company = *request.Company
	}
	if request.Locate != nil {
		opening.Locate = *request.Locate
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != nil {
		opening.Link = *request.Link
	}
	if request.Salary != nil {
		opening.Salary = *request.Salary
	}
	if request.Description != nil {
		opening.Description = *request.Description
	}

	// 3. Salva usando o Service
	updatedOpening, err := h.service.UpdateOpening(opening)
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "Update opening error")
		return
	}

	handler.SendSuccess(ctx, "update-opening", updatedOpening)
}
