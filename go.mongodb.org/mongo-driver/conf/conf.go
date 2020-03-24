package conf

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

func getCfg() *ini.File {
	cfg, err := ini.Load("./go.mongodb.org/mongo-driver/conf/app.ini")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}

func GetMongo() *ini.Section {
	cfg := getCfg()
	db := cfg.Section("database")
	return db
}
