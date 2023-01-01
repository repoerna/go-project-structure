package user

import (
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

const TABLE_NAME = "users"

type Column struct {
	UID,
	Username,
	Email,
	FullName,
	HashedPassword,
	CreatedAt,
	UpdatedAt string
}

var Columns = Column{
	UID:            "uid",
	Username:       "username",
	Email:          "email",
	FullName:       "full_name",
	HashedPassword: "hashed_password",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

func (c Column) ManipulatedColumns() []string {
	return []string{
		c.UID,
		c.Username,
		c.Email,
		c.FullName,
		c.HashedPassword,
		c.CreatedAt,
		c.UpdatedAt,
	}
}

type User struct {
	UID            uuid.UUID
	Username       string
	Email          string
	FullName       string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Users []User

func (u *User) MockUser() {
	u.UID = uuid.New()
	u.Username = faker.Username()
	u.Email = faker.Email()
	u.FullName = faker.Name()
	u.HashedPassword = faker.Password()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (us *Users) MockUsers(count int) {
	for count > 0 {
		var user User
		user.MockUser()
		*us = append(*us, user)
		count--
	}
}
