package mysql

import (
	"context"
	"fmt"
	"gorm-v2/config"
	"gorm-v2/util"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       connectionString,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mysql db")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(100)

	if cfg.Debug {
		db = db.Debug()
	}
}

// nolint
func GetClient(ctx context.Context) *gorm.DB {
	cloneDB := &gorm.DB{}
	*cloneDB = *db

	// use User per request
	if util.IsEnableTx(ctx) {
		tx := util.GetTx(ctx)
		return tx
	}

	return cloneDB
}

func Disconnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if db != nil {
		sqlDB.Close()
	}
}
