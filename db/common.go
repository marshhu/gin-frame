package db

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)

var dbMapPool map[string]*sql.DB
var gormMapPool map[string]*gorm.DB

func GetDBPool(name string) (*sql.DB, error) {
	if dbPool, ok := dbMapPool[name]; ok {
		return dbPool, nil
	}
	return nil, fmt.Errorf("GetDBPool %s  error", name)
}

func GetGormPool(name string) (*gorm.DB, error) {
	if dbPool, ok := gormMapPool[name]; ok {
		return dbPool, nil
	}
	return nil, fmt.Errorf("GetGormPool %s  error", name)
}

type Settings struct {
	DriverName      string `json:"driver-name"`
	DbName          string `json:"db-name"`
	DataSourceName  string `json:"data-source-name"`
	MaxOpenConn     int    `json:"max-open-conn"`
	MaxIdleConn     int    `json:"max-idle-conn"`
	MaxConnLifeTime int    `json:"max-conn-life-time"`
}
