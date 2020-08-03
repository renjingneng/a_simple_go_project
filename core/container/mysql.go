package container

import (
	"database/sql"

	//this is
	_ "github.com/go-sql-driver/mysql"

	"github.com/renjingneng/a_simple_go_project/core/config"
)

var mysqlContainer map[string]*sql.DB

//GetEntityFromMysqlContainer is
func GetEntityFromMysqlContainer(database string, mode string) *sql.DB {
	if database == "" || mode == "" {
		return nil
	}
	dbname := database + mode
	if db, ok := mysqlContainer[dbname]; ok {
		return db
	}
	if _, ok := config.DatabaseMap[dbname]; !ok {
		return nil
	}
	if db, err := sql.Open("mysql", config.DatabaseMap[dbname]); err != nil {
		return nil
	} else {
		mysqlContainer[dbname] = db
		return db
	}
}
func init() {
	if mysqlContainer == nil {
		mysqlContainer = make(map[string]*sql.DB)
	}
}
