package mid

import (
	"context"
	"net/http"

	"github.com/St3plox/Blogchain/foundation/web"
)

func Cors() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")

			return handler(ctx, w, r)

		}

		return h
	}

	return m
}
