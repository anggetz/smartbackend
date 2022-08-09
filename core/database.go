package core

import (
	"fmt"
	"smartservice/helpers"
	"smartservice/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) GetConnection() (*Database, error) {
	var config helpers.Config
	var err error

	err = helpers.ReadConfig("./config/config.yaml", &config)
	if err != nil {
		return nil, err
	}

	dsn := config.Database
	d.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Database) GormQuery() *gorm.DB {
	return d.db
}

func (d *Database) Close() error {
	if d.db != nil {
		sqlDB, err := d.db.DB()
		if err != nil {
			return err
		}

		err = sqlDB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) MigrationModel() error {
	_, err := d.GetConnection()
	if err != nil {
		return err
	}
	defer d.Close()

	if d.db == nil {
		return fmt.Errorf("please connect to database first")
	}

	d.db.Migrator().AutoMigrate(&model.User{})
	d.db.Migrator().AutoMigrate(&model.Blok{})
	d.db.Migrator().AutoMigrate(&model.GroupLocation{})

	return nil
}
