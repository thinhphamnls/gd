package gdrepo

import (
	"errors"
	"github.com/thinhphamnls/gd/container"
	"github.com/thinhphamnls/gd/model/gd"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IQuickBooksOAuthRepo interface {
	List(param *gdmodel.QuickBooksAuthQuery, schema string) ([]*gdmodel.QuickBooksAuthView, bool, error)
}

type qbOAuthRepo struct {
	db    container.IDataBaseProvider
	sugar *zap.SugaredLogger
}

func NewQuickBooksOAuthRepo(db container.IDataBaseProvider, sugar *zap.SugaredLogger) IQuickBooksOAuthRepo {
	return &qbOAuthRepo{
		db:    db,
		sugar: sugar,
	}
}

func (r *qbOAuthRepo) List(param *gdmodel.QuickBooksAuthQuery, schema string) ([]*gdmodel.QuickBooksAuthView, bool, error) {
	var (
		tx     = r.db.GDSlave()
		result []*gdmodel.QuickBooksAuthView
	)

	if param.AppUserName != "" {
		tx = tx.Where("qo.app_username = ?", param.AppUserName)
	}

	if param.AppTenant != "" {
		tx.Where("qo.app_tenant = ?", param.AppTenant)
	}

	if err := tx.
		Table(r.tableName(schema)).
		Where("qo.deleted = ?", param.Deleted).
		Find(&result).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, err
	}

	return result, len(result) > 0, nil
}

func (r *qbOAuthRepo) tableName(schema string) string {
	return buildTableName(schema, "quickbooks_oauth", "qo")
}
