package account

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi/operations"
)

// userCreate is the handler of the user creation endpoint.
// This func calls the user creation endpoint of account-svc with the given data.
func (h *RestHandler) userCreate(params operations.UserCreateParams) middleware.Responder {
	// Call endpoint to create a new user with the given input.
	resp, err := h.accountService.CreateUser(params.HTTPRequest.Context(), &accountproto.CreateUserRequest{
		User: &accountproto.User{
			Name: params.Seed.Name,
		},
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

	// Convert proto model to the Swagger model.
	model := toUserModel(resp.GetUser())

	// Return the created user model.
	return operations.NewUserCreateOK().WithPayload(model)
}
