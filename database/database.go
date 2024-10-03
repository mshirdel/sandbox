package database

import (
	"fmt"

	"github.com/mshirdel/sandbox/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SandboxDatabase struct {
	DB *gorm.DB
}

func NewSandboxDatabase() *SandboxDatabase {
	return &SandboxDatabase{}
}

func (d *SandboxDatabase) Init() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sandbox?charset=utf8mb4&parseTime=True&loc=Local"

	dialect := mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error in opening database: %w", err)
	}
	
	d.DB = db

	if err = db.AutoMigrate(&models.User{}, &models.CreditCard{}); err != nil {
		return fmt.Errorf("error in auto migration: %w", err)
	}
	
	return nil
}

func (d *SandboxDatabase) Close() error {

	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
