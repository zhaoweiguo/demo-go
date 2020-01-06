package main

import (
	"database/sql"
	"github.com/drone/drone/store/shared/migrate/sqlite"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func main() {
	driver := "sqlite3"
	datasource := "core.sqlite"
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db, err := sql.Open(driver, datasource)
	if err != nil {
		log.Panic(err)
	}
	if err := pingDatabase(db); err != nil {
		log.Panic(err)
	}
	if err := sqlite.Migrate(db); err != nil {
		log.Panic(err)
	}

	sqlxdb := sqlx.NewDb(db, driver)

}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}
