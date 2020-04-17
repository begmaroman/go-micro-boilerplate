package account

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi/operations"
)

// userDelete is the handler of the user deletion endpoint.
// This func calls the user deletion endpoint of account-svc with the given data.
func (h *RestHandler) userDelete(params operations.UserDeleteParams) middleware.Responder {
	// Call endpoint to delete an existing user with the given ID.
	resp, err := h.accountService.DeleteUser(params.HTTPRequest.Context(), &accountproto.DeleteUserRequest{
		UserId: params.UserID.String(),
	})
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

	// Return nothing, just 204 status code.
	return operations.NewUserDeleteNoContent()
}
