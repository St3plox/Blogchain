package likegrp

import v1 "github.com/St3plox/Blogchain/app/backend/user-service/handlers/v1"

const basePath = "/like"

const GetPath = v1.Version + basePath + "/" + "{id}"

const GetByUserOrPostIDPath = v1.Version + basePath

const PostPath = v1.Version + basePath

const PutPath = v1.Version + basePath

const DeletePath = v1.Version + basePath + "/" + "{id}"
