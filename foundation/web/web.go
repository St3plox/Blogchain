package web

import (
	"context"
	"errors"
	"net/http"
	"os"
	"syscall"
	"time"

	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*mux.Router
	shutdown chan os.Signal
	mw       []Middleware
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		Router:   mux.NewRouter(),
		shutdown: shutdown,
		mw:       mw,
	}
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(path string, method string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		var v = Values{
			TraceID: uuid.NewString(),
			Now:     time.Now().UTC(),
		}
		ctx = context.WithValue(ctx, key, &v)

		if err := handler(ctx, w, r); err != nil {
			if validateShutdown(err) {
				a.SignalShutdown()
				return nil // Return nil here to indicate graceful shutdown without error
			}
			return err
		}
		return nil // Return nil here to indicate successful handling without error
	}

	// Wrap 'h' in http.HandlerFunc to match the ServeMux.Handler method signature
	a.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		err := h(r.Context(), w, r)
		if err != nil {

			var reqErr *v1.RequestError
			if errors.As(err, &reqErr) {

				http.Error(w, reqErr.Err.Error(), reqErr.Status)
				return
			}

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}).Methods(method)
}
func validateShutdown(err error) bool {

	// Ignore syscall.EPIPE and syscall.ECONNRESET errors which occurs
	// when a write operation happens on the http.ResponseWriter that
	// has simultaneously been disconnected by the client (TCP
	// connections is broken). For instance, when large amounts of
	// data is being written or streamed to the client.
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://gosamples.dev/broken-pipe/
	// https://gosamples.dev/connection-reset-by-peer/

	switch {
	case errors.Is(err, syscall.EPIPE):

		// Usually, you get the broken pipe error when you write to the connection after the
		// RST (TCP RST Flag) is sent.
		// The broken pipe is a TCP/IP error occurring when you write to a stream where the
		// other end (the peer) has closed the underlying connection. The first write to the
		// closed connection causes the peer to reply with an RST packet indicating that the
		// connection should be terminated immediately. The second write to the socket that
		// has already received the RST causes the broken pipe error.
		return false

	case errors.Is(err, syscall.ECONNRESET):

		// Usually, you get connection reset by peer error when you read from the
		// connection after the RST (TCP RST Flag) is sent.
		// The connection reset by peer is a TCP/IP error that occurs when the other end (peer)
		// has unexpectedly closed the connection. It happens when you send a packet from your
		// end, but the other end crashes and forcibly closes the connection with the RST
		// packet instead of the TCP FIN, which is used to close a connection under normal
		// circumstances.
		return false
	}

	return true
}
