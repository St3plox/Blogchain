package mid

import (
	"context"
	"net/http"

	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/foundation/web"
)

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate(a *auth.Auth) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			claims, err := a.Authenticate(ctx, r.Header.Get("authorization"))
			if err != nil {
				return auth.NewAuthError("authenticate: failed: %s", err)
			}

			ctx = auth.SetClaims(ctx, claims)

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}

// Authorize validates that an authenticated user has at least one role from a
// specified list. This method constructs the actual function that is used.
func Authorize(a *auth.Auth, rule string) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			claims := auth.GetClaims(ctx)

			if claims.Subject == "" {
				return auth.NewAuthError("authorize: you are not authorized for that action, no claims")
			}

			if err := a.Authorize(ctx, claims, rule); err != nil {
				return auth.NewAuthError("authorize: you are not authorized for that action, claims[%v] rule[%v]: %s", claims.Roles, rule, err)
			}

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
