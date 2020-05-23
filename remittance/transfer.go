package remittance

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	transfer "github.com/wondenge/momo-go/remittance/gen/transfer"
	"goa.design/goa/v3/security"
)

// transfer service example implementation.
// The example methods log the requests and return zero values.
type transfersrvc struct {
	logger log.Logger
}

// NewTransfer returns the transfer service implementation.
func NewTransfer(logger log.Logger) transfer.Service {
	return &transfersrvc{logger}
}

// JWTAuth implements the authorization logic for service "transfer" for the
// "jwt" security scheme.
func (s *transfersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// Transfers an amount from the owner’s account to a payee account
func (s *transfersrvc) Create(ctx context.Context, p *transfer.CreatePayload) (res *transfer.Errorreason, err error) {
	res = &transfer.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("transfer.create"))
	return
}

// This operation is used to get the status of a transfer.
func (s *transfersrvc) Get(ctx context.Context, p *transfer.GetPayload) (res *transfer.Errorreason, err error) {
	res = &transfer.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("transfer.get"))
	return
}
