package account

import (
	"github.com/go-openapi/strfmt"
	"github.com/golang/protobuf/ptypes"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/models"
)

// toUserModel converts the user proto model to the Swagger model.
func toUserModel(u *accountproto.User) *models.User {
	updatedAt, _ := ptypes.Timestamp(u.GetUpdatedAt())
	createdAt, _ := ptypes.Timestamp(u.GetCreatedAt())

	return &models.User{
		ID:        u.GetId(),
		Name:      u.GetName(),
		UpdatedAt: strfmt.DateTime(updatedAt),
		CreatedAt: strfmt.DateTime(createdAt),
	}
}
