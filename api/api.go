package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nuts-foundation/nuts-pgo-demo/domain/accounts"
	"github.com/nuts-foundation/nuts-pgo-demo/domain/types"
	"net/http"
)

type Wrapper struct {
	Auth           auth
	AccountService accounts.Service
}

func (w Wrapper) CreateAccount(ctx echo.Context) error {
	request := types.CreateAccountRequest{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	_, err := w.AccountService.CreateAccount(request.Username, request.Password)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (w Wrapper) CreateSession(ctx echo.Context) error {
	sessionRequest := types.CreateSessionRequest{}
	if err := ctx.Bind(&sessionRequest); err != nil {
		return err
	}

	if !w.Auth.CheckCredentials(sessionRequest.Username, sessionRequest.Password) {
		return echo.NewHTTPError(http.StatusForbidden, "invalid credentials")
	}

	token, err := w.Auth.CreateJWT(sessionRequest.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, types.CreateSessionResponse{Token: string(token)})
}

func (w Wrapper) CheckSession(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
