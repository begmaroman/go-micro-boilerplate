package account

import (
	"time"

	"github.com/go-openapi/strfmt"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/models"
)

// toUserModel converts the user proto model to the Swagger model.
func toUserModel(u *accountproto.User) *models.User {
	var updatedAt strfmt.DateTime
	if updatedAtRaw := u.GetUpdatedAt(); updatedAtRaw > 0 {
		updatedAt = strfmt.DateTime(time.Unix(int64(updatedAtRaw), 0))
	}

	var createdAt strfmt.DateTime
	if createdAtRaw := u.GetCreatedAt(); createdAtRaw > 0 {
		createdAt = strfmt.DateTime(time.Unix(int64(createdAtRaw), 0))
	}

	return &models.User{
		ID:        u.GetId(),
		Name:      u.GetName(),
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
}
