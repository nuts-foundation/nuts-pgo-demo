package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nuts-foundation/go-did/did"
	"github.com/nuts-foundation/nuts-pgo-demo/domain/types"
	"net/http"
)

type Wrapper struct {
	Auth auth
}

func (w Wrapper) CreateAccount(ctx echo.Context) error {
	request := types.CreateAccountRequest{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	// use a fake DID to prevent parsing errors. Must be replaced with a real one.
	fakeDID := did.MustParseDID("did:nuts:123")
	newAccount := types.Account{Username: request.Username, Password: request.Password, DID: fakeDID}
	_, err := w.Auth.accounts.AddAccount(newAccount)
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
