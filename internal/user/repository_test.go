package user

import (
	"context"
	"database/sql"
	"fmt"
	"go-boilerplate-api/pkg/pagination"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = Describe("User Repository", func() {
	var repo Repository
	var mock sqlmock.Sqlmock
	var ctx context.Context = context.Background()

	var users Users
	users.MockUsers(10)

	BeforeEach(func() {
		var db *sql.DB
		var err error

		// mock sql.DB
		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		// mock open gorm db
		gdb, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}))
		Expect(err).ShouldNot(HaveOccurred())

		repo = Repository{db: gdb}

	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("Get", func() {
		type args struct {
			ctx  context.Context
			page *pagination.PageOffset
		}
		tests := []struct {
			name    string
			args    args
			want    Users
			want1   int64
			wantErr bool
		}{
			// TODO: Add test cases.
			// {
			// 	"return empty",
			// 	args{
			// 		ctx,
			// 		&pagination.PageOffset{
			// 			Offset: 0,
			// 			Limit:  0,
			// 		},
			// 	},
			// 	Users{},
			// 	0,
			// 	false,
			// },
			{
				"return data offset 0 limit 2",
				args{
					ctx,
					&pagination.PageOffset{
						Offset: 0,
						Limit:  2,
					},
				},
				Users{users[0], users[1]},
				2,
				false,
			},
		}
		for _, tt := range tests {
			It(tt.name, func() {
				const sqlSelectAll = `SELECT *`
				const sqlCount = `SELECT count`

				var rows *sqlmock.Rows = sqlmock.NewRows(Columns.ManipulatedColumns())
				if len(tt.want) > 0 {
					for _, usr := range tt.want {
						rows.AddRow(
							usr.UID,
							usr.Username,
							usr.Email,
							usr.FullName,
							usr.HashedPassword,
							usr.CreatedAt,
							usr.UpdatedAt,
						)
					}
				}

				mock.ExpectQuery(sqlCount).WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(tt.want1))

				mock.ExpectQuery(sqlSelectAll).
					WillReturnRows(rows)

				got, got1, err := repo.Get(tt.args.ctx, tt.args.page)
				fmt.Println(err)

				if (err != nil) != tt.wantErr {
					Expect(err).Should(Equal(tt.wantErr))
					return
				}

				Expect(got).Should(Equal(tt.want))
				Expect(got1).Should(Equal(tt.want1))

			})
		}
	})

})

// func TestNewRepository(t *testing.T) {
// 	type args struct {
// 		db *gorm.DB
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want Repository
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRepository_Get(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ctx  context.Context
// 		page *pagination.Pagination
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    Users
// 		want1   *pagination.TotalEntries
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := Repository{
// 				db: tt.fields.db,
// 			}
// 			got, got1, err := r.Get(tt.args.ctx, tt.args.page)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Repository.Get() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("Repository.Get() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// func TestRepository_GetByID(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ctx context.Context
// 		id  uuid.UUID
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *User
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := Repository{
// 				db: tt.fields.db,
// 			}
// 			got, err := r.GetByID(tt.args.ctx, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Repository.GetByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRepository_Insert(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ctx  context.Context
// 		user *User
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *User
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := Repository{
// 				db: tt.fields.db,
// 			}
// 			got, err := r.Insert(tt.args.ctx, tt.args.user)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.Insert() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Repository.Insert() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRepository_Update(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ctx  context.Context
// 		user *User
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *User
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := Repository{
// 				db: tt.fields.db,
// 			}
// 			got, err := r.Update(tt.args.ctx, tt.args.user)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.Update() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Repository.Update() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRepository_Delete(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ctx context.Context
// 		id  uuid.UUID
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := Repository{
// 				db: tt.fields.db,
// 			}
// 			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestRepository_Seed(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ctx   context.Context
// 		count int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := Repository{
// 				db: tt.fields.db,
// 			}
// 			r.Seed(tt.args.ctx, tt.args.count)
// 		})
// 	}
// }
