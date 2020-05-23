package disbursement

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	token "github.com/wondenge/momo-go/disbursement/gen/token"
	"goa.design/goa/v3/security"
)

// token service example implementation.
// The example methods log the requests and return zero values.
type tokensrvc struct {
	logger log.Logger
}

// NewToken returns the token service implementation.
func NewToken(logger log.Logger) token.Service {
	return &tokensrvc{logger}
}

// BasicAuth implements the authorization logic for service "token" for the
// "basic" security scheme.
func (s *tokensrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
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

// Creates an Access Token.
func (s *tokensrvc) Create(ctx context.Context, p *token.CreatePayload) (res *token.OAuthTokenError, err error) {
	res = &token.OAuthTokenError{}
	s.logger.Log("info", fmt.Sprintf("token.create"))
	return
}
