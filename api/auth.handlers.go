package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	router "github.com/thedivinez/go-libs/gothex"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (*AuthHandler) HandleSignUp(ctx echo.Context) error {
	/*
		user := types.User{
			Email:    ctx.FormValue("email"),
			Password: ctx.FormValue("password"),
			Username: ctx.FormValue("username"),
		}
		add user to database logic here
	*/
	log.Println(ctx.FormValue("email"))
	return ctx.String(http.StatusOK, "incorrect username or password")
}

func (*AuthHandler) HandleSignIn(ctx echo.Context) error {
	if ctx.QueryParam("provider") != "" {
		if user, err := gothic.CompleteUserAuth(ctx.Response().Writer, ctx.Request()); err == nil {
			log.Println(user)
		} else {
			gothic.BeginAuthHandler(ctx.Response().Writer, ctx.Request())
		}
		return nil
	}
	if router.IsHxRequest(ctx) {
		return router.SignIn(ctx, "/", map[string]any{"name": "Divine Zvenyika"})
	}
	return ctx.String(http.StatusOK, "logged in")
}

func (*AuthHandler) HandleSignOut(ctx echo.Context) error {
	if ctx.QueryParam("provider") != "" {
		gothic.Logout(ctx.Response().Writer, ctx.Request())
	}
	if router.IsHxRequest(ctx) {
		return router.SignOut(ctx, "/signin")
	}
	return ctx.String(http.StatusOK, "logged out")
}

func (*AuthHandler) HandleProvidersCallback(ctx echo.Context) error {
	user, err := gothic.CompleteUserAuth(ctx.Response().Writer, ctx.Request())
	if err != nil {
		return err
	}
	log.Println(user)
	return nil
}
