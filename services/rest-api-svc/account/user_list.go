package account

import (
	"net/http"

	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/models"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi/operations"
)

// usersList is the handler of the users listing endpoint.
// This func calls the users listing endpoint of account-svc.
func (h *RestHandler) usersList(params operations.UsersListParams) middleware.Responder {
	// Call endpoint to list all users.
	resp, err := h.accountService.ListUsers(params.HTTPRequest.Context(), &accountproto.ListUsersRequest{})
	if err != nil {
		// Handle the given error and return 500 status code.
		// Also, write error message into the response.
		// Error handling can be with more clear way, but it's just an example.
		return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		})
	} else if resp.GetError().GetCode() != 0 {
		// Handle the given logic error and return 400 status code.
		// Also, write error message into the response.
		// Error handling can be with more clear way, but it's just an example.
		// Here gonna be mapping between RPC and HTTP status codes.
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			http.Error(w, resp.GetError().GetMessage(), http.StatusBadRequest)
		})
	}

	// Convert proto models to the Swagger models.
	users := make([]*models.User, len(resp.GetData().GetUsers()))
	for i, user := range resp.GetData().GetUsers() {
		users[i] = toUserModel(user)
	}

	// Return user models.
	return operations.NewUsersListOK().WithPayload(users)
}
