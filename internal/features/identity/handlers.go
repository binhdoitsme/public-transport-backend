package identity

import (
	"net/http"
	commonErrors "public-transport-backend/internal/common/errors"
	"public-transport-backend/internal/common/responses"
	"public-transport-backend/internal/features/identity/createtokens"
	"public-transport-backend/internal/features/identity/invalidatetokens"
	"public-transport-backend/internal/features/identity/me"
	"public-transport-backend/internal/features/identity/refreshtokens"
	"public-transport-backend/internal/features/identity/signup"
	"time"

	"github.com/gin-gonic/gin"
)

type Dependencies interface {
	CreateTokenPairDependenciesFactory() *createtokens.Dependencies
	RefreshTokenPairDependenciesFactory() *refreshtokens.Dependencies
	InvalidateTokenPairDependenciesFactory() *invalidatetokens.Dependencies
	SignUpDependenciesFactory() *signup.Dependencies
	GetMyProfileDependenciesFactory() *me.Dependencies
}

type handler struct {
	dependencies Dependencies
}

func (h *handler) handleNewProfile(ctx *gin.Context) {
	form := &signup.SignUpForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}

	result, err := signup.CreateUserAccount(ctx, form, h.dependencies.SignUpDependenciesFactory())
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.Data(ctx, http.StatusCreated, result)
}

func (h *handler) handleGetMyProfile(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.NotAuthorizedError().Error())
		return
	}
	form := &me.GetMyProfileForm{UserId: userId.(uint64)}

	result, err := me.GetMyProfile(ctx, form, h.dependencies.GetMyProfileDependenciesFactory())
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.Data(ctx, http.StatusOK, result)
}

func (h *handler) handleNewTokenPair(ctx *gin.Context) {
	form := &createtokens.NewTokensForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}
	result, err := createtokens.NewTokenPair(ctx, form, h.dependencies.CreateTokenPairDependenciesFactory())
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if !result.Ok {
		responses.Data(ctx, http.StatusUnauthorized, result)
		return
	}
	responses.Data(ctx, http.StatusCreated, result)
}

func (h *handler) handleRefreshTokenPair(ctx *gin.Context) {
	form := &refreshtokens.RefreshTokenForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}
	form.Now = time.Now().UTC()
	result, err := refreshtokens.RefreshTokenPair(
		ctx, form, h.dependencies.RefreshTokenPairDependenciesFactory())

	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if !result.Ok {
		responses.Data(ctx, http.StatusUnauthorized, result)
		return
	}
	responses.Data(ctx, http.StatusCreated, result)
}

func (h *handler) handleInvalidateTokenPair(ctx *gin.Context) {
	form := &invalidatetokens.InvalidateTokenForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}
	form.Now = time.Now().UTC()
	result, err := invalidatetokens.InvalidateToken(
		ctx, form, h.dependencies.InvalidateTokenPairDependenciesFactory())

	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.Data(ctx, http.StatusOK, result)
}

func (h *handler) handleChangePassword(ctx *gin.Context) {}

func InitAPIHandlers(g *gin.RouterGroup, dependencies Dependencies) {
	h := &handler{dependencies}
	tokens := g.Group("/v1/tokens")
	{
		tokens.POST("/", h.handleNewTokenPair)
		tokens.POST("/refresh", h.handleRefreshTokenPair)
		tokens.DELETE("/", h.handleInvalidateTokenPair)
	}
	profile := g.Group("/v1/profile")
	{
		profile.POST("/", h.handleNewProfile)
		profile.PATCH("/password", h.handleChangePassword)
		profile.POST("/me", h.handleGetMyProfile)
	}
}

func InitMiddlewares(g *gin.RouterGroup, dependencies Dependencies) {
	g.Use(func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		account, err := dependencies.CreateTokenPairDependenciesFactory().Tokens.Parse(token)
		if err != nil || account == nil {
			return
		}
		ctx.Set("userId", account.Id)
	})
}
