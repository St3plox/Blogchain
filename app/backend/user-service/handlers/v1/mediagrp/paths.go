package mediagrp

import v1 "github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1"

const BasePath = "/media"

const PostPath = v1.Version + BasePath

const PostMultiple = v1.Version + BasePath + "/many"

const GetPath = v1.Version + BasePath + "/{" + MediaID + "}"

const GetByIDsPath = v1.Version + BasePath

const DeletePath = v1.Version + BasePath + "/{" + MediaID + "}"

const MediaID = "mediaID"
