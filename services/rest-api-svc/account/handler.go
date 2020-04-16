package account

import (
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi/operations"
)

// RestHandlerOptions contains required options for the handler.
// This handler implements REST endpoints with handling incoming data.
type RestHandlerOptions struct {
	AccountService accountproto.AccountService
	Logger         logrus.FieldLogger
}

// RestHandler defines the REST interface for the business service.
type RestHandler struct {
	accountService accountproto.AccountService
	logger         logrus.FieldLogger
}

// NewRestHandler creates a new Handler.
func NewRestHandler(opts *RestHandlerOptions) *RestHandler {
	return &RestHandler{
		accountService: opts.AccountService,
		logger:         opts.Logger,
	}
}

// Register registers endpoints to the handler.
func (h *RestHandler) Register(api *operations.RestAPISvcAPI) {
	api.UserCreateHandler = operations.UserCreateHandlerFunc(h.userCreate)
	api.UserReadHandler = operations.UserReadHandlerFunc(h.userRead)
	api.UsersListHandler = operations.UsersListHandlerFunc(h.usersList)
	api.UserUpdateHandler = operations.UserUpdateHandlerFunc(h.userUpdate)
	api.UserDeleteHandler = operations.UserDeleteHandlerFunc(h.userDelete)
}
