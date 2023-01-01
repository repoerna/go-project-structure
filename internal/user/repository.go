package user

import (
	"context"
	"go-boilerplate-api/pkg/pagination"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// Repository implements IRepository

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Get(ctx context.Context, page *pagination.PageOffset) (Users, int64, error) {
	var users []User
	var count int64

	if err := r.db.Model(&User{}).Count(&count).Error; err != nil {
		return nil, count, err
	}

	if page != nil {
		if err := r.db.Limit(page.Limit).Offset(page.Offset).Find(&users).Error; err != nil {
			return nil, count, err
		}
	}

	return users, count, nil
}

func (r Repository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return nil, nil
}

func (r Repository) Insert(ctx context.Context, user *User) (*User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r Repository) Update(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (r Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r Repository) Seed(ctx context.Context, count int) {
	for count > 0 {
		user := User{
			UID:            uuid.New(),
			Username:       faker.Username(),
			Email:          faker.Email(),
			FullName:       faker.Name(),
			HashedPassword: faker.Password(),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		r.Insert(ctx, &user)
		count--
	}
}
