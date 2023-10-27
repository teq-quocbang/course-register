package auth

import (
	"fmt"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/util/myerror"
	"github.com/teq-quocbang/course-register/util/token"
)

type key string

const (
	UserPrincipleKey key = "user-principle"
	AuthorizationKey key = "x-server-auth-key"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		t := ctx.Request().Header.Get(string(AuthorizationKey))
		if t == "" {
			return teq.Response.Error(ctx, myerror.ErrForbidden(fmt.Errorf("missing token")))
		}

		configs := config.GetConfig()
		jwt := token.JWT{
			SecretKey: configs.TokenSecretKey,
		}

		claims, err := jwt.VerifyToken(t)
		if err != nil {
			return teq.Response.Error(ctx, myerror.ErrForbidden(err))
		}

		echo.Context.Set(ctx, string(UserPrincipleKey), claims)

		return next(ctx)
	}
}
