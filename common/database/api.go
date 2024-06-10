package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"team-management/common/config"
	"time"
)

func Init(cfg *config.Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DbDSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	//check whether database was initialised
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(20)
	return db, nil
}
