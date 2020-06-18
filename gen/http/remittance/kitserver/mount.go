// Code generated by goa v3.1.3, DO NOT EDIT.
//
// remittance go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountNewTokenHandler configures the mux to serve the "remittance" service
// "NewToken" endpoint.
func MountNewTokenHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/remittance/token", f)
}

// MountGetBalanceHandler configures the mux to serve the "remittance" service
// "GetBalance" endpoint.
func MountGetBalanceHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/remittance/v1_0/account/balance", f)
}

// MountRetrieveAccountHolderHandler configures the mux to serve the
// "remittance" service "RetrieveAccountHolder" endpoint.
func MountRetrieveAccountHolderHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/remittance/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active", f)
}

// MountTransferHandler configures the mux to serve the "remittance" service
// "Transfer" endpoint.
func MountTransferHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/remittance/v1_0/transfer", f)
}

// MountTransferStatusHandler configures the mux to serve the "remittance"
// service "TransferStatus" endpoint.
func MountTransferStatusHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/remittance/v1_0/transfer/{referenceId}", f)
}
