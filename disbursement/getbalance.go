package disbursement

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	getbalance "github.com/wondenge/momo-go/disbursement/gen/getbalance"
	"goa.design/goa/v3/security"
)

// getbalance service example implementation.
// The example methods log the requests and return zero values.
type getbalancesrvc struct {
	logger log.Logger
}

// NewGetbalance returns the getbalance service implementation.
func NewGetbalance(logger log.Logger) getbalance.Service {
	return &getbalancesrvc{logger}
}

// JWTAuth implements the authorization logic for service "getbalance" for the
// "jwt" security scheme.
func (s *getbalancesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// Get the balance of the account
func (s *getbalancesrvc) Show(ctx context.Context, p *getbalance.ShowPayload) (res *getbalance.Errorreason, err error) {
	res = &getbalance.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("getbalance.show"))
	return
}
