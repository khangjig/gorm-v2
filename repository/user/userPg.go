package user

import (
	"context"
	"gorm-v2/model"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return pgUserRepository{getDB}
}

type pgUserRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p pgUserRepository) GetTableName(ctx context.Context) string {
	stmt := &gorm.Statement{DB: p.getDB(ctx)}

	err := stmt.Parse(&model.User{})
	if err != nil {
		log.Fatal("get table name has fail")
	}

	return stmt.Schema.Table
}

func (p pgUserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var data model.User
	err := p.getDB(ctx).First(&data, id).Error

	return &data, errors.Wrap(err, "GetByID fail")
}

func (p pgUserRepository) GetAll(ctx context.Context,
	conditions []model.Condition,
	paginator model.Paginator,
	order []string,
) ([]model.User, int64) {
	db := p.getDB(ctx)

	for _, condition := range conditions {
		switch strings.ToLower(condition.Type) {
		case "not":
			db = db.Not(condition.Pattern, condition.Values...)
		case "or":
			db = db.Or(condition.Pattern, condition.Values...)
		default:
			db = db.Where(condition.Pattern, condition.Values...)
		}
	}

	for i := range order {
		db = db.Order(order[i])
	}

	var response []model.User

	total := db.Find(&response).RowsAffected

	err := db.Limit(paginator.Limit).Offset(paginator.Page).Find(&response).Error
	if err != nil {
		// Handle log ...
	}

	return response, total
}

func (p pgUserRepository) Create(ctx context.Context, obj *model.User) error {
	tx := p.getDB(ctx).Begin()

	if err := tx.Create(obj).Error; err != nil {
		tx.Rollback()

		return errors.Wrap(err, "Create fail")
	}

	err := tx.Commit().Error

	return errors.Wrap(err, "Commit fail")
}

func (p pgUserRepository) ExecuteSQL(ctx context.Context, query string) error {
	return p.getDB(ctx).Exec(query).Error
}
