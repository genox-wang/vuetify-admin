package model

import (
	log "github.com/sirupsen/logrus"

	"vuetify-admin-api/app/config"

	"github.com/jinzhu/gorm"
)

var (
	// DB global gorm instance
	DB *gorm.DB
)

// OpenDB open grom
func OpenDB(db string) {
	var err error
	log.Info("OpeningDB: ", db)
	DB, err = gorm.Open("mysql", db)

	DB.LogMode(config.GetBool("db.showlog"))

	if err != nil {
		log.Fatalf(err.Error())
	}
	Migrate()
	createAdminUser()
}

// Migrate db migration
func Migrate() {
	log.Info("Migrate ... ")
	DB.AutoMigrate(
		new(User),
		new(Channel),
		new(Server),
	)
}

// CloseDB close gorm
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func createAdminUser() {
	u := &User{
		DisplayName: "admin",
		Username:    "admin",
		Password:    "admin",
	}
	u.Create()
}
