// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountShowHandler configures the mux to serve the "getbalance" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/remittance/v1_0/account/balance", f)
}