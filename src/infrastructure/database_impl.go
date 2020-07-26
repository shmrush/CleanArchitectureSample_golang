package infrastructure

import (
	"log"

	"CleanArchitectureSample_golang/common"

	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database provides gorm connection.
type DatabaseImpl struct {
	Conn *gorm.DB
}

// NewDatabase initializes Database.
func NewDatabase() Database {
	env := common.DatabaseEnv
	db, err := gorm.Open(env.Kind, env.User+":"+env.Password+"@tcp("+env.Host+")/"+env.DbName+"?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)
	if err != nil {
		log.Fatal(err.Error())
	}
	return Database{db}
}
