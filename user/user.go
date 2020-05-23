package userapi

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	user "github.com/wondenge/momo-go/user/gen/user"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger log.Logger) user.Service {
	return &usersrvc{logger}
}

// Used to create an API user in the sandbox target environment
func (s *usersrvc) Createuser(ctx context.Context, p *user.CreateuserPayload) (res *user.Errorreason, err error) {
	res = &user.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("user.createuser"))
	return
}

// Used to create an API key for an API user in the sandbox target environment.
func (s *usersrvc) Createkey(ctx context.Context, p *user.CreatekeyPayload) (res *user.Errorreason, err error) {
	res = &user.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("user.createkey"))
	return
}

// Used to get API user information.
func (s *usersrvc) List(ctx context.Context, p *user.ListPayload) (res *user.Errorreason, err error) {
	res = &user.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("user.list"))
	return
}

// GET API User Details
func (s *usersrvc) Show(ctx context.Context, p *user.ShowPayload) (res *user.Errorreason, err error) {
	res = &user.Errorreason{}
	s.logger.Log("info", fmt.Sprintf("user.show"))
	return
}
