package remittance

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	validate "github.com/wondenge/momo-go/remittance/gen/validate"
	"goa.design/goa/v3/security"
)

// validate service example implementation.
// The example methods log the requests and return zero values.
type validatesrvc struct {
	logger log.Logger
}

// NewValidate returns the validate service implementation.
func NewValidate(logger log.Logger) validate.Service {
	return &validatesrvc{logger}
}

// JWTAuth implements the authorization logic for service "validate" for the
// "jwt" security scheme.
func (s *validatesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// Checks if an account holder is registered and active in the system
func (s *validatesrvc) Show(ctx context.Context, p *validate.ShowPayload) (res *validate.Errorreason, err error) {
	res = &validate.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("validate.show"))
	return
}
