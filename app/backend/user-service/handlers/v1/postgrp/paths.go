package postgrp

import v1 "github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1"

const BasePath = "/posts"

const PostPath = v1.Version + BasePath

const GetUserPostsPath = v1.Version + BasePath

const GetAllPath = v1.Version + BasePath + "/all"

const GetPostsByUserAddressPath = v1.Version + BasePath + "/{address}"

const GetByIdPath = v1.Version + BasePath + "/id/{id}"

const GetPostsByUserAddressAndIndexPath = v1.Version + BasePath + "/{address}/{index}"
