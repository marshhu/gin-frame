package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMysql(settings []Settings) error {
	for _, dbSetting := range settings {
		dbPool, err := sql.Open("mysql", dbSetting.DataSourceName)
		if err != nil {
			return err
		}
		dbPool.SetMaxOpenConns(dbSetting.MaxOpenConn)
		dbPool.SetMaxIdleConns(dbSetting.MaxIdleConn)
		dbPool.SetConnMaxLifetime(time.Duration(dbSetting.MaxConnLifeTime) * time.Second)
		err = dbPool.Ping()
		if err != nil {
			return err
		}

		//gorm连接方式
		dbGorm, err := gorm.Open(mysql.New(mysql.Config{Conn: dbPool}))
		if err != nil {
			return err
		}
		if dbMapPool == nil {
			dbMapPool = make(map[string]*sql.DB)
		}
		if gormMapPool == nil {
			gormMapPool = make(map[string]*gorm.DB)
		}
		dbMapPool[dbSetting.DbName] = dbPool
		gormMapPool[dbSetting.DbName] = dbGorm
	}
	return nil
}
