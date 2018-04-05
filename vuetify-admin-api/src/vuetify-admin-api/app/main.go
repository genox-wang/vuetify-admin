package main

import (
	_ "github.com/go-sql-driver/mysql"

	"vuetify-admin-api/app/config"
	"vuetify-admin-api/app/model"
	"vuetify-admin-api/app/router"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLogLevel()
	log.Warnf("DB => %s", config.GetString("db.mysql"))
	model.OpenDB(config.GetString("db.mysql"))
	defer model.CloseDB()
	router.Run(8000)
}

func initLogLevel() {
	switch config.GetString("log.logLevel") {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	}
}
