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
	dbConnMaxLifetime = 60 * time.Minute
)

const (
	DefaultSchemaGorillaDesk = "gorilladesk"
)

type DataBaseProvider interface {
	Main() *gorm.DB
	Slave() *gorm.DB
	Transaction(fc func(tx *gorm.DB) error) (err error)
}

type databaseProvider struct {
	main, slave *gorm.DB
}

func NewDatabase(cfMain gdconfig.DbConfig, cfSlave gdconfig.DbConfig, zap gdlogger.ZapLoggerProvider) (DataBaseProvider, func(), error) {
	var (
		data = &databaseProvider{}
		err  error
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

		closeDB(data.main)
		closeDB(data.slave)

		zap.Get().Info("closing the db repo resources")
	}

	if cfMain.Host != "" {
		data.main, err = connect(cfMain, zap)
		if err != nil {
			return data, cleanup, err
		}
	}

	if cfSlave.Host != "" {
		data.slave, err = connect(cfSlave, zap)
		if err != nil {
			return data, cleanup, err
		}
	}

	return data, cleanup, nil
}

func connect(cf gdconfig.DbConfig, zap gdlogger.ZapLoggerProvider) (*gorm.DB, error) {
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

func (p *databaseProvider) Main() *gorm.DB {
	return p.main.Session(&gorm.Session{SkipHooks: false})
}

func (p *databaseProvider) Slave() *gorm.DB {
	return p.slave.Session(&gorm.Session{SkipHooks: false})
}

func (p *databaseProvider) Transaction(fc func(tx *gorm.DB) error) (err error) {
	panicked := true
	tx := p.main.Begin()

	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txRep := p.main
	txRep = tx
	err = fc(txRep)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
