package handlers

import (
	"fmt"
	"net/http"
	"profotdel-rest/internal/modules/category/v1/application/dto"
	"profotdel-rest/internal/modules/category/v1/application/mapper"
	"profotdel-rest/internal/modules/category/v1/application/usecase"
	"profotdel-rest/internal/shared/lib/request"
	"profotdel-rest/internal/shared/lib/response"
	"profotdel-rest/internal/shared/lib/validator"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	createUC usecase.CreateUseCase
	getAllUC usecase.GetAllUseCase
}

func NewHandler(createUC usecase.CreateUseCase, getAllUC usecase.GetAllUseCase) Handler {
	return &handler{
		createUC: createUC,
		getAllUC: getAllUC,
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	createDTO, err := request.DecodeBody[dto.CreateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := validator.ValidateDTO(createDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	model := mapper.ToModelFromCreateDTO(&createDTO)
	createdModel, errFields, err := h.createUC.Execute(ctx, model)
	if err != nil {
		msg := fmt.Sprintf("failed to create category: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	responseDTO := mapper.ToResponseDTOFromModel(createdModel)

	response.SendSuccess(w, http.StatusCreated, responseDTO)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	models, err := h.getAllUC.Execute(ctx)
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve categories: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	responseDTOs := make([]*dto.ResponseDTO, len(models))
	for i, model := range models {
		responseDTOs[i] = mapper.ToResponseDTOFromModel(model)
	}

	response.SendSuccess(w, http.StatusOK, responseDTOs)
}
