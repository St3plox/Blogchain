package usergrp

import v1 "github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1"

const BasePath = "/users"

const RegisterUserPath = v1.Version + BasePath + "/register"

const LoginUserPath = v1.Version + BasePath + "/login"
