package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nuts-foundation/nuts-pgo-demo/domain"
	"net/http"
)

type Wrapper struct {
	Auth auth
}

func (w Wrapper) CreateSession(ctx echo.Context) error {
	sessionRequest := domain.CreateSessionRequest{}
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

	return ctx.JSON(http.StatusOK, domain.CreateSessionResponse{Token: string(token)})
}

func (w Wrapper) CheckSession(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
