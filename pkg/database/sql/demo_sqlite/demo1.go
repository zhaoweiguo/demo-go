package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	driver := "sqlite3"
	datasource := "./pkg/database/sql/demo_sqlite/core.sqlite"
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 创建数据库，生成core.sqlite文件
	db, err := sql.Open(driver, datasource)
	if err != nil {
		log.Panic(err)
	}

	// 测试连通性
	if err := pingDatabase(db); err != nil {
		log.Panic(err)
	}

	// 创建migration表
	if _, err := db.Exec(migrationTableCreate); err != nil {
		log.Panic(err)
	}

	// 创建子表,也即user与repo表,并往migration表插入2条数据
	for _, migration := range migrations {
		if _, err := db.Exec(migration.stmt); err != nil {
			log.Panic(err)
		}
		if _, err := db.Exec(migrationInsert, migration.name); err != nil {
			log.Panic(err)
		}

	}

	// 查询
	rows, err := db.Query(migrationSelect)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Panic(err)
		}
		log.Println(name)
	}
	//
	//sqlxdb := sqlx.NewDb(db, driver)

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

var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
 name VARCHAR(255)
)
`
var migrationInsert = `
INSERT INTO migrations (name) VALUES (?)
`

var migrationSelect = `
SELECT name FROM migrations
`

var migrations = []struct {
	name string
	stmt string
}{
	{
		name: "create-table-users",
		stmt: createTableUsers,
	},
	{
		name: "create-table-repos",
		stmt: createTableRepos,
	},
}

var createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
);
`

//
// 002_create_table_repos.sql
//

var createTableRepos = `
CREATE TABLE IF NOT EXISTS repos (
 repo_id                    INTEGER PRIMARY KEY AUTOINCREMENT
,repo_uid                   TEXT
,repo_user_id               INTEGER
,repo_namespace             TEXT
,UNIQUE(repo_uid)
);
`
