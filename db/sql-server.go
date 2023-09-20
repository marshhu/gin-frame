package db

import (
	"database/sql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

func InitSqlServer(settings []Settings) error {
	for _, dbSetting := range settings {
		dbPool, err := sql.Open("sqlserver", dbSetting.DataSourceName)
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
		dbGorm, err := gorm.Open(sqlserver.Open(dbSetting.DataSourceName))
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
