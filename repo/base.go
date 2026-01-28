package gdrepo

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/thinhphamnls/gd/container"
)

type IQuery interface {
	BuildQuery(tx *gorm.DB) *gorm.DB
}

type BaseRepo[T any] struct {
	db    gdcontainer.DataBaseProvider
	sugar *zap.SugaredLogger
}

func NewBaseRepo[T any](db gdcontainer.DataBaseProvider, sugar *zap.SugaredLogger) *BaseRepo[T] {
	return &BaseRepo[T]{
		db:    db,
		sugar: sugar,
	}
}

func (r *BaseRepo[T]) List(param IQuery, table string) ([]*T, bool, error) {
	var (
		tx     = r.db.Slave()
		result []*T
	)

	tx = param.BuildQuery(tx)

	if err := tx.
		Debug().
		Table(table).
		Find(&result).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, err
	}

	return result, len(result) > 0, nil
}

func (r *BaseRepo[T]) Get(param IQuery, table string) (*T, bool, error) {
	var (
		tx     = r.db.Slave()
		result *T
	)

	tx = param.BuildQuery(tx)

	if err := tx.Table(table).
		First(&result).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, err
	}

	return result, result != nil, nil
}

func (r *BaseRepo[T]) Edit(tx *gorm.DB, param IQuery, updates map[string]interface{}, table string) error {
	tx = param.BuildQuery(tx)
	return tx.Table(table).
		Updates(updates).
		Error
}

func (r *BaseRepo[T]) Create(tx *gorm.DB, input interface{}, table string) error {
	return tx.Omit(clause.Associations).
		Table(table).
		Create(input).Error
}

func (r *BaseRepo[T]) Delete(tx *gorm.DB, ent interface{}, ids []uint, table string) error {
	return tx.
		Table(table).
		Delete(ent, ids).Error
}
