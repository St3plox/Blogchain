package handlers

import (
	"os"

	"github.com/St3plox/Blogchain/app/backend/user-service/handlers/postgrp"
	"github.com/St3plox/Blogchain/app/backend/user-service/handlers/usergrp"
	"github.com/St3plox/Blogchain/business/core/post"
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
	PostCore *post.Core
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics(), mid.Cors())

	uh := usergrp.New(cfg.UserCore, cfg.Auth)

	app.Handle("/users/register", "POST", uh.RegisterUser)
	app.Handle("/users/login", "POST", uh.LoginUser)

	ph := postgrp.New(cfg.PostCore, cfg.Auth, cfg.UserCore)

	app.Handle("/posts", "POST", ph.Post, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))

	app.Handle("/posts", "GET", ph.GetUserPosts, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))
	app.Handle("/posts/all", "GET", ph.GetAll)
	app.Handle("/posts/{address}", "GET", ph.GetPostsByUserAddress, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))

	app.Handle("/posts/id/{id}", "GET", ph.GetById, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))
	app.Handle("/posts/{address}/{index}", "GET", ph.GetPostsByUserAddressAndIndex, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))

	return app
}
