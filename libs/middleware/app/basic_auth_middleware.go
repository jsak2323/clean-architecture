package app

import (
	errs "clean-architecture/libs/errors"
	"clean-architecture/libs/httpresponse"
	"clean-architecture/libs/try"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func (cfg *CfgMiddlerware) VerifyBasiAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		headerAuth := ctx.Request().Header.Get("Authorization")
		basicAuthToken := try.ArrayStringToString(strings.Split(headerAuth, "Basic "), 1, "")
		auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("BASIC_AUTH_USERNAME") + ":" + os.Getenv("BASIC_AUTH_PASSWORD")))
		if auth != basicAuthToken {
			e := errors.New(errs.InvalidAuthentication)
			httpresponse.Error(ctx, http.StatusUnauthorized, httpresponse.ErrMessage{
				ErrMapping: e,
				Message:    errs.ErrBasicAuthMSG,
			})
			return e
		}

		return next(ctx)
	}
}
