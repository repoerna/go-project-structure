package user

import (
	"context"
	"go-boilerplate-api/pkg/pagination"

	"github.com/google/uuid"
)

type IRepository interface {
	Get(context.Context, *pagination.PageOffset) (Users, int64, error)
	GetByID(context.Context, uuid.UUID) (*User, error)
	Insert(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	Delete(context.Context, uuid.UUID) error
}
