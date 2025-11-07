package gdcontainer

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/thinhphamnls/gd/config"
	"github.com/thinhphamnls/gd/logger"
)

const (
	dbConnMaxLifetime = 3 * time.Minute
)

const (
	DefaultSchemaGorillaDesk = "gorilladesk"
)

type IDataBaseProvider interface {
	GDMain() *gorm.DB
	GDSlave() *gorm.DB
	Transaction(fc func(tx *gorm.DB) error) (err error)
}

type databaseProvider struct {
	gdMain, gdSlave *gorm.DB
}

func NewDatabase(cf gdconfig.IBaseConfig, zap gdlogger.IBaseLogger) (IDataBaseProvider, func(), error) {
	var (
		data = &databaseProvider{}

		cfGDMain  = cf.GetDatabase().GDMain
		cfGDSlave = cf.GetDatabase().GDSlave

		err error
	)

	cleanup := func() {
		closeDB := func(db *gorm.DB) {
			if db != nil {
				sqlDB, err := db.DB()
				if err == nil {
					_ = sqlDB.Close()
				}
			}
		}

		closeDB(data.gdMain)
		closeDB(data.gdSlave)

		zap.Get().Info("closing the db repo resources")
	}

	if cfGDMain.Host != "" {
		data.gdMain, err = connect(cfGDMain, zap)
		if err != nil {
			return data, cleanup, err
		}
	}

	if cfGDSlave.Host != "" {
		data.gdSlave, err = connect(cfGDSlave, zap)
		if err != nil {
			return data, cleanup, err
		}
	}

	return data, cleanup, nil
}

func connect(cf gdconfig.DbConfig, zap gdlogger.IBaseLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cf.Host, cf.Username, cf.Password, cf.DBName, cf.Port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{Logger: zap})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cf.MaxIdleCon)
	sqlDB.SetMaxOpenConns(cf.MaxCon)
	sqlDB.SetConnMaxLifetime(dbConnMaxLifetime)

	if sqlDB == nil {
		return nil, errors.New("cannot open connection to database")
	}

	return db, nil
}

func (p *databaseProvider) GDMain() *gorm.DB {
	return p.gdMain.Session(&gorm.Session{SkipHooks: false})
}

func (p *databaseProvider) GDSlave() *gorm.DB {
	return p.gdSlave.Session(&gorm.Session{SkipHooks: false})
}

func (p *databaseProvider) Transaction(fc func(tx *gorm.DB) error) (err error) {
	panicked := true
	tx := p.gdMain.Begin()

	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txRep := p.gdMain
	txRep = tx
	err = fc(txRep)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
