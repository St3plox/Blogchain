package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/rs/zerolog"
)

func Logger(log *zerolog.Logger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v := web.GetValues(ctx)

			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			log.Info().
				Str("trace_id", v.TraceID).
				Str("method", r.Method).
				Str("path", path).
				Str("remoteaddr", r.RemoteAddr).
				Msg("request started")

			err := handler(ctx, w, r)

			log.Info().
				Str("trace_id", v.TraceID).
				Str("method", r.Method).
				Str("path", path).
				Str("remoteaddr", r.RemoteAddr).
				Int("statuscode", v.StatusCode).
				Dur("since", time.Since(v.Now)).
				Msg("request completed")

			return err
		}

		return h
	}

	return m
}
