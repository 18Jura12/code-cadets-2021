package bootstrap

import (
	"database/sql"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
)

func Sqlite() *sql.DB {
	db, err := sql.Open("sqlite3", config.Cfg.SqliteDatabase)
	if err != nil {
		panic(err)
	}

	return db
}
