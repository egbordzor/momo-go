package collection

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	requestopay "github.com/wondenge/momo-go/collection/gen/requestopay"
	"goa.design/goa/v3/security"
)

// requestopay service example implementation.
// The example methods log the requests and return zero values.
type requestopaysrvc struct {
	logger log.Logger
}

// NewRequestopay returns the requestopay service implementation.
func NewRequestopay(logger log.Logger) requestopay.Service {
	return &requestopaysrvc{logger}
}

// JWTAuth implements the authorization logic for service "requestopay" for the
// "jwt" security scheme.
func (s *requestopaysrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// Request a payment from a consumer (Payer).
func (s *requestopaysrvc) Create(ctx context.Context, p *requestopay.CreatePayload) (res *requestopay.Errorreason, err error) {
	res = &requestopay.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("requestopay.create"))
	return
}

// Get the status of a request to pay.
func (s *requestopaysrvc) Get(ctx context.Context, p *requestopay.GetPayload) (res *requestopay.Errorreason, err error) {
	res = &requestopay.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("requestopay.get"))
	return
}
