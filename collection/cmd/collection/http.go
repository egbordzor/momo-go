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
	getbalance "github.com/wondenge/momo-go/collection/gen/getbalance"
	getbalancekitsvr "github.com/wondenge/momo-go/collection/gen/http/getbalance/kitserver"
	getbalancesvr "github.com/wondenge/momo-go/collection/gen/http/getbalance/server"
	requestopaykitsvr "github.com/wondenge/momo-go/collection/gen/http/requestopay/kitserver"
	requestopaysvr "github.com/wondenge/momo-go/collection/gen/http/requestopay/server"
	tokenkitsvr "github.com/wondenge/momo-go/collection/gen/http/token/kitserver"
	tokensvr "github.com/wondenge/momo-go/collection/gen/http/token/server"
	validatekitsvr "github.com/wondenge/momo-go/collection/gen/http/validate/kitserver"
	validatesvr "github.com/wondenge/momo-go/collection/gen/http/validate/server"
	requestopay "github.com/wondenge/momo-go/collection/gen/requestopay"
	token "github.com/wondenge/momo-go/collection/gen/token"
	validate "github.com/wondenge/momo-go/collection/gen/validate"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, tokenEndpoints *token.Endpoints, validateEndpoints *validate.Endpoints, getbalanceEndpoints *getbalance.Endpoints, requestopayEndpoints *requestopay.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

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
		tokenCreateHandler       *kithttp.Server
		tokenServer              *tokensvr.Server
		validateShowHandler      *kithttp.Server
		validateServer           *validatesvr.Server
		getbalanceShowHandler    *kithttp.Server
		getbalanceServer         *getbalancesvr.Server
		requestopayCreateHandler *kithttp.Server
		requestopayGetHandler    *kithttp.Server
		requestopayServer        *requestopaysvr.Server
	)
	{
		eh := errorHandler(logger)
		tokenCreateHandler = kithttp.NewServer(
			endpoint.Endpoint(tokenEndpoints.Create),
			tokenkitsvr.DecodeCreateRequest(mux, dec),
			tokenkitsvr.EncodeCreateResponse(enc),
			kithttp.ServerErrorEncoder(tokenkitsvr.EncodeCreateError(enc, nil)),
		)
		tokenServer = tokensvr.New(tokenEndpoints, mux, dec, enc, eh, nil)
		validateShowHandler = kithttp.NewServer(
			endpoint.Endpoint(validateEndpoints.Show),
			validatekitsvr.DecodeShowRequest(mux, dec),
			validatekitsvr.EncodeShowResponse(enc),
			kithttp.ServerErrorEncoder(validatekitsvr.EncodeShowError(enc, nil)),
		)
		validateServer = validatesvr.New(validateEndpoints, mux, dec, enc, eh, nil)
		getbalanceShowHandler = kithttp.NewServer(
			endpoint.Endpoint(getbalanceEndpoints.Show),
			getbalancekitsvr.DecodeShowRequest(mux, dec),
			getbalancekitsvr.EncodeShowResponse(enc),
			kithttp.ServerErrorEncoder(getbalancekitsvr.EncodeShowError(enc, nil)),
		)
		getbalanceServer = getbalancesvr.New(getbalanceEndpoints, mux, dec, enc, eh, nil)
		requestopayCreateHandler = kithttp.NewServer(
			endpoint.Endpoint(requestopayEndpoints.Create),
			requestopaykitsvr.DecodeCreateRequest(mux, dec),
			requestopaykitsvr.EncodeCreateResponse(enc),
			kithttp.ServerErrorEncoder(requestopaykitsvr.EncodeCreateError(enc, nil)),
		)
		requestopayGetHandler = kithttp.NewServer(
			endpoint.Endpoint(requestopayEndpoints.Get),
			requestopaykitsvr.DecodeGetRequest(mux, dec),
			requestopaykitsvr.EncodeGetResponse(enc),
			kithttp.ServerErrorEncoder(requestopaykitsvr.EncodeGetError(enc, nil)),
		)
		requestopayServer = requestopaysvr.New(requestopayEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	tokenkitsvr.MountCreateHandler(mux, tokenCreateHandler)
	validatekitsvr.MountShowHandler(mux, validateShowHandler)
	getbalancekitsvr.MountShowHandler(mux, getbalanceShowHandler)
	requestopaykitsvr.MountCreateHandler(mux, requestopayCreateHandler)
	requestopaykitsvr.MountGetHandler(mux, requestopayGetHandler)

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
	for _, m := range tokenServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range validateServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range getbalanceServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range requestopayServer.Mounts {
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
