package passenger

import (
	"net/http"
	commonErrors "public-transport-backend/internal/common/errors"
	"public-transport-backend/internal/common/responses"
	"public-transport-backend/internal/features/passenger/create"
	"public-transport-backend/internal/features/passenger/view"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Dependencies interface {
	CreateDependenciesFactory() *create.Dependencies
	ViewDependenciesFactory() *view.Dependencies
}

type handler struct {
	dependencies Dependencies
}

func (h *handler) handleSelfCreatePassenger(ctx *gin.Context) {
	var form *create.SelfPassengerForm = &create.SelfPassengerForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}
	dependencies := h.dependencies.CreateDependenciesFactory()
	result, err := create.SelfCreatePassenger(ctx, form, dependencies)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.Data(ctx, http.StatusCreated, result)
}

func (h *handler) handleAdminCreatePassenger(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		responses.Error(ctx, http.StatusUnauthorized, commonErrors.NotAnAdminError().Error())
		return
	}

	var form *create.AdminPassengerForm = &create.AdminPassengerForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}
	form.SetAdminUserId(userId.(uint64))

	dependencies := h.dependencies.CreateDependenciesFactory()

	result, err := create.AdminCreatePassenger(ctx, form, dependencies)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.Data(ctx, http.StatusCreated, result)
}

func (h *handler) handleListPassengers(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		responses.Error(ctx, http.StatusUnauthorized, commonErrors.NotAnAdminError().Error())
		return
	}
	requestingUser := &view.RequestingUser{
		UserId: userId.(uint64),
	}

	page := 1
	pageSize := 50

	if pageParam, exists := ctx.Params.Get("page"); exists {
		page, _ = strconv.Atoi(pageParam)
	}

	if pageSizeParam, exists := ctx.Params.Get("pagesize"); exists {
		pageSize, _ = strconv.Atoi(pageSizeParam)
	}
	form := &view.PassengerListForm{
		RequestingUser: requestingUser,
		Page:           page,
		PageSize:       pageSize,
	}

	dependencies := h.dependencies.ViewDependenciesFactory()

	passengers, err := view.AdminListPassengers(ctx, form, dependencies)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.Data(ctx, http.StatusOK, passengers)
}

func (h *handler) handleViewMyPassenger(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		responses.Error(ctx, http.StatusUnauthorized, commonErrors.NotAuthorizedError().Error())
		return
	}

	requestingUser := &view.RequestingUser{
		UserId: userId.(uint64),
	}

	dependencies := h.dependencies.ViewDependenciesFactory()

	passenger, err := view.ViewMyPassenger(ctx, requestingUser, dependencies)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.Data(ctx, http.StatusOK, passenger)
}

func (h *handler) handleAdminViewOnePassenger(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		responses.Error(ctx, http.StatusUnauthorized, commonErrors.NotAnAdminError().Error())
		return
	}
	requestingUser := &view.RequestingUser{
		UserId: userId.(uint64),
	}
	passengerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, commonErrors.ToGenericError(err).Error())
		return
	}

	dependencies := h.dependencies.ViewDependenciesFactory()

	passengers, err := view.AdminViewPassenger(ctx, &view.AdminPassengerByIdForm{
		Id:             uint64(passengerId),
		RequestingUser: requestingUser,
	}, dependencies)

	if err != nil {
		var statusCode int
		if strings.HasPrefix(err.Error(), "ERR_013") {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusBadRequest
		}
		responses.Error(ctx, statusCode, err.Error())
		return
	}

	responses.Data(ctx, http.StatusOK, passengers)
}

// func (h *handler) handlePassengerApproval(ctx *gin.Context) {}

func InitAPIHandlers(g *gin.RouterGroup, dependencies Dependencies) {
	h := &handler{dependencies}
	api := g.Group("/v1/passengers")
	{
		api.POST("/", h.handleSelfCreatePassenger)
		api.POST("/admin", h.handleAdminCreatePassenger)

		api.GET("/", h.handleListPassengers)
		api.GET("/me", h.handleViewMyPassenger)
		api.GET("/:id", h.handleAdminViewOnePassenger)
		// api.GET("/{id}/history", h.handleViewOnePassengerHistory)

		// api.POST("/{id}/approval", h.handlePassengerApproval)
		// api.PATCH("/{id}/type", h.handleChangeAccountType)
		// api.DELETE("/{id}", h.handleDeleteOnePassenger)
	}
}
