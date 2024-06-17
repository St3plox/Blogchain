package handlers

import (
	"os"

	"github.com/St3plox/Blogchain/app/backend/user-service/handlers/maingrp"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/business/web/v1/mid"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/rs/zerolog"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zerolog.Logger
	Auth     *auth.Auth
	UserCore *user.Core
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics())

	h := maingrp.New(cfg.UserCore, cfg.Auth)

	app.Handle("GET /sucker", h.Get, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))
	app.Handle("POST /register", h.RegisterUser)
	app.Handle("POST /login", h.LoginUser)

	return app
}
