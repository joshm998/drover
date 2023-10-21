package db

import (
	"github.com/joshm998/drover/config"
	"github.com/joshm998/drover/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqlCon struct {
	DBPool *gorm.DB
}

var conn *sqlCon

type sqlConn struct {
	DbPool *gorm.DB
}

var connector *sqlConn

func InitMysql() *sqlConn {
	if connector != nil {
		log.Info("DataBase is initialized")
		return connector
	}
	log.Info("DataBase was not initialized ..initializing again")
	var err error
	connector, err = initDB()
	if err != nil {
		panic(err)
	}
	return connector
}

// DB Initialization

func initDB() (*sqlConn, error) {

	dbFile := config.GetYamlValues().DBConfig.FileName

	db, err := gorm.Open(sqlite.Open(dbFile))
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&model.Printers{})
	// if maxCons := config.GetYamlValues().DBConfig.MaxConnection; maxCons > 0 {
	// 	db.DB().SetMaxOpenConns(maxCons)
	// 	db.DB().SetMaxIdleConns(maxCons / 3)
	// }
	return &sqlConn{db}, nil
}

func GetDBConnection() *gorm.DB {
	return connector.DbPool
}
