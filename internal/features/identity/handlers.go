package identity

import (
	"net/http"
	commonErrors "public-transport-backend/internal/common/errors"
	"public-transport-backend/internal/common/responses"
	"public-transport-backend/internal/features/identity/createtokens"
	"public-transport-backend/internal/features/identity/signup"

	"github.com/gin-gonic/gin"
)

type Dependencies interface {
	LoginDependenciesFactory() *createtokens.Dependencies
	SignUpDependenciesFactory() *signup.Dependencies
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

}

func (h *handler) handleNewTokenPair(ctx *gin.Context) {
	form := &createtokens.NewTokensForm{}
	if err := ctx.BindJSON(form); err != nil {
		responses.Error(ctx, http.StatusBadRequest, commonErrors.ToValidationError(err).Error())
		return
	}
	result, err := createtokens.NewSession(ctx, form, h.dependencies.LoginDependenciesFactory())
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if !result.Ok {
		responses.Data(ctx, http.StatusUnauthorized, result)
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    result.AccessToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	})
	responses.Data(ctx, http.StatusCreated, result)
}

func (h *handler) handleRefreshTokenPair(ctx *gin.Context) {}
func (h *handler) handleRevokeTokenPair(ctx *gin.Context)  {}
func (h *handler) handleChangePassword(ctx *gin.Context)   {}

func InitAPIHandlers(g *gin.RouterGroup, dependencies Dependencies) {
	h := &handler{dependencies}
	tokens := g.Group("/v1/tokens")
	{
		tokens.POST("/", h.handleNewTokenPair)
		tokens.POST("/refresh", h.handleRefreshTokenPair)
		tokens.DELETE("/", h.handleRevokeTokenPair)
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
		token, err := ctx.Cookie("token")
		if err != nil {
			return
		}
		account, err := dependencies.LoginDependenciesFactory().Tokens.Parse(token)
		if err != nil {
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "token",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			})
			return
		}

		// set user info from token
		if account == nil {
			return
		}
		ctx.Set("userId", account.Id)
	})
}
