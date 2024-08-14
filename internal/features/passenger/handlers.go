package passenger

import (
	"net/http"
	commonErrors "public-transport-backend/internal/common/errors"
	"public-transport-backend/internal/common/responses"
	"public-transport-backend/internal/features/passenger/create"

	"github.com/gin-gonic/gin"
)

type Dependencies interface {
	CreateDependenciesFactory() *create.Dependencies
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
	res, err := create.SelfCreatePassenger(ctx, form, dependencies)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	asJson := struct {
		Data *create.CreatePassengerResult `json:"data"`
	}{res}
	ctx.JSON(http.StatusCreated, asJson)
}

func (h *handler) handleAdminCreatePassenger(ctx *gin.Context) {
	var form *create.AdminPassengerForm = &create.AdminPassengerForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}

	form.UserId = ctx.GetUint64("userId")

	dependencies := h.dependencies.CreateDependenciesFactory()

	res, err := create.AdminCreatePassenger(ctx, form, dependencies)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	asJson := struct {
		Data *create.CreatePassengerResult `json:"data"`
	}{res}
	ctx.JSON(http.StatusCreated, asJson)
}

// func (h *handler) handlePassengerApproval(ctx *gin.Context) {

// }

func InitAPIHandlers(g *gin.RouterGroup, dependencies Dependencies) {
	h := &handler{dependencies}
	api := g.Group("/v1/passengers")
	{
		api.POST("/", h.handleSelfCreatePassenger)
		api.POST("/admin", h.handleAdminCreatePassenger)

		// api.GET("/", h.handleListPassengers)
		// api.GET("/{id}", h.handleViewOnePassenger)
		// api.GET("/{id}/history", h.handleViewOnePassengerHistory)

		// api.POST("/{id}/approval", h.handlePassengerApproval)
		// api.PATCH("/{id}/type", h.handleChangeAccountType)
		// api.DELETE("/{id}", h.handleDeleteOnePassenger)
	}
}
