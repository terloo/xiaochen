package db

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {

	var err error
	DB, err = gorm.Open(sqlite.Open(C.Path), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("failed to connect database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(C.MaxIdleConns)
	sqlDB.SetMaxOpenConns(C.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(C.ConnMaxLifetime) * time.Second)

}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
