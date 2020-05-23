package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	userkitsvr "github.com/wondenge/momo-go/user/gen/http/user/kitserver"
	usersvr "github.com/wondenge/momo-go/user/gen/http/user/server"
	user "github.com/wondenge/momo-go/user/gen/user"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, userEndpoints *user.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		userCreateuserHandler *kithttp.Server
		userCreatekeyHandler  *kithttp.Server
		userListHandler       *kithttp.Server
		userShowHandler       *kithttp.Server
		userServer            *usersvr.Server
	)
	{
		eh := errorHandler(logger)
		userCreateuserHandler = kithttp.NewServer(
			endpoint.Endpoint(userEndpoints.Createuser),
			userkitsvr.DecodeCreateuserRequest(mux, dec),
			userkitsvr.EncodeCreateuserResponse(enc),
			kithttp.ServerErrorEncoder(userkitsvr.EncodeCreateuserError(enc, nil)),
		)
		userCreatekeyHandler = kithttp.NewServer(
			endpoint.Endpoint(userEndpoints.Createkey),
			userkitsvr.DecodeCreatekeyRequest(mux, dec),
			userkitsvr.EncodeCreatekeyResponse(enc),
			kithttp.ServerErrorEncoder(userkitsvr.EncodeCreatekeyError(enc, nil)),
		)
		userListHandler = kithttp.NewServer(
			endpoint.Endpoint(userEndpoints.List),
			userkitsvr.DecodeListRequest(mux, dec),
			userkitsvr.EncodeListResponse(enc),
			kithttp.ServerErrorEncoder(userkitsvr.EncodeListError(enc, nil)),
		)
		userShowHandler = kithttp.NewServer(
			endpoint.Endpoint(userEndpoints.Show),
			userkitsvr.DecodeShowRequest(mux, dec),
			userkitsvr.EncodeShowResponse(enc),
			kithttp.ServerErrorEncoder(userkitsvr.EncodeShowError(enc, nil)),
		)
		userServer = usersvr.New(userEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	userkitsvr.MountCreateuserHandler(mux, userCreateuserHandler)
	userkitsvr.MountCreatekeyHandler(mux, userCreatekeyHandler)
	userkitsvr.MountListHandler(mux, userListHandler)
	userkitsvr.MountShowHandler(mux, userShowHandler)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(logger)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range userServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
