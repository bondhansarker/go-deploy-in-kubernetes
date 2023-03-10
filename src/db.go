package src

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dB  *gorm.DB
	err error
)

func DbClient() *gorm.DB {
	return dB
}

func ConnectDb(dbConfig *DbConfig) {
	logMode := logger.Silent
	if dbConfig.Debug {
		logMode = logger.Info
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Schema)
	dB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode),
	})
	if err != nil {
		panic(err)
	}

	sqlDb, err := dB.DB()
	if err != nil {
		panic(err)
	}

	if dbConfig.MaxIdleConn != 0 {
		sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConn)
	}
	if dbConfig.MaxOpenConn != 0 {
		sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConn)
	}
	if dbConfig.MaxConnLifetime != 0 {
		sqlDb.SetConnMaxLifetime(dbConfig.MaxConnLifetime * time.Second)
	}
	dB.AutoMigrate(
	// &User{},
	)
	fmt.Println("mysql connection successful...")
}
