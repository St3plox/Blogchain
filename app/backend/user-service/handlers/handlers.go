package handlers

import (
	"os"

	"github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1/mediagrp"
	"github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1/postgrp"
	"github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1/usergrp"
	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/St3plox/Blogchain/business/core/post"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/business/web/v1/mid"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/rs/zerolog"
)

type APIMuxConfig struct {
	Shutdown  chan os.Signal
	Log       *zerolog.Logger
	Auth      *auth.Auth
	UserCore  *user.Core
	PostCore  *post.Core
	MediaCore *media.Core
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics(), mid.Cors())

	uh := usergrp.New(cfg.UserCore, cfg.Auth)

	app.Handle(usergrp.RegisterUserPath, "POST", uh.RegisterUser)
	app.Handle(usergrp.LoginUserPath, "POST", uh.LoginUser)

	ph := postgrp.New(cfg.PostCore, cfg.Auth, cfg.UserCore)

	app.Handle(postgrp.PostPath, "POST", ph.Post, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))

	app.Handle(postgrp.GetUserPostsPath, "GET", ph.GetUserPosts, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAny))

	app.Handle(postgrp.GetAllPath, "GET", ph.GetAll)
	app.Handle(postgrp.GetPostsByUserAddressPath, "GET", ph.GetPostsByUserAddress)
	app.Handle(postgrp.GetByIdPath, "GET", ph.GetById)
	app.Handle(postgrp.GetPostsByUserAddressAndIndexPath, "GET", ph.GetPostsByUserAddressAndIndex)

	mc := mediagrp.New(cfg.MediaCore, cfg.UserCore)

	app.Handle(mediagrp.GetPath, "GET", mc.Get)
	app.Handle(mediagrp.PostPath, "POST", mc.Post)
	app.Handle(mediagrp.UpdatePath, "PUT", mc.Update)
	app.Handle(mediagrp.DeletePath, "DELETE", mc.Delete)
 
	return app
}
