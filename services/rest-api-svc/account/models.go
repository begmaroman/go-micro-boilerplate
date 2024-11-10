package account

import (
	"github.com/go-openapi/strfmt"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/models"
)

// toUserModel converts the user proto model to the Swagger model.
func toUserModel(u *accountproto.User) *models.User {
	var updatedAt strfmt.DateTime
	if updatedAtRaw := u.GetUpdatedAt(); updatedAtRaw != nil {
		updatedAt = strfmt.DateTime(u.GetUpdatedAt().AsTime())
	}

	var createdAt strfmt.DateTime
	if createdAtRaw := u.GetCreatedAt(); createdAtRaw != nil {
		createdAt = strfmt.DateTime(u.GetCreatedAt().AsTime())
	}

	return &models.User{
		ID:        u.GetId(),
		Name:      u.GetName(),
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
}
