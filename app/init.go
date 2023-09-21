package app

import (
	"github.com/marshhu/gin-frame/common"
	"github.com/marshhu/gin-frame/config"
	"github.com/marshhu/gin-frame/db"
	"github.com/marshhu/gin-frame/log"
	"github.com/marshhu/gin-frame/redis"
	"github.com/marshhu/gin-frame/utils"
	"github.com/spf13/viper"
	sysLog "log"
	"os"
	"strings"
)

// InitModules 模块初始化
func InitModules(modules []string) error {
	// 初始化配置
	err := InitConf()
	if err != nil {
		return err
	}
	// 初始化日志
	logSetting := GetLogSettings()
	log.Init(&logSetting)
	// 初始化sql-server
	if utils.StringInArray(common.SqlServerModule, modules) {
		dbSettings := GetDbSettings(common.SqlServerModule)
		if err = db.InitSqlServer(dbSettings); err != nil {
			return err
		}
	}
	// 初始化mysql
	if utils.StringInArray(common.MysqlModule, modules) {
		dbSettings := GetDbSettings(common.MysqlModule)
		if err = db.InitSqlServer(dbSettings); err != nil {
			return err
		}
	}
	// 初始化redis
	if utils.StringInArray(common.RedisModule, modules) {
		redisSettings := GetRedisSettings()
		if err = redis.InitClient(redisSettings); err != nil {
			return err
		}
	}
	return nil
}

func InitConf() error {
	useLocalConf := false
	useRemoteConf := strings.TrimSpace(os.Getenv(common.UseRemoteConf))
	if strings.ToLower(useRemoteConf) == "true" {
		err := config.InitRemote()
		if err != nil {
			sysLog.Printf("加载远程配置失败：%v", err)
			useLocalConf = true
		} else {
			sysLog.Println("当前使用的配置是：远程配置")
			return nil
		}
	} else {
		useLocalConf = true
	}
	if useLocalConf {
		err := config.InitLocal()
		if err != nil {
			sysLog.Printf("加载本地配置失败：%v", err)
			return err
		}
	}
	sysLog.Println("当前使用的配置是：本地配置")
	return nil
}

func GetLogSettings() log.Settings {
	settings := log.Settings{
		Path:        viper.GetString("log.path"),
		FileName:    viper.GetString("log.file-name"),
		Level:       viper.GetString("log.level"),
		LogCategory: viper.GetString("log.category"),
		Caller:      true,
	}
	return settings
}

func GetDbSettings(confNodeName string) []db.Settings {
	sqlServerArray := viper.Get(confNodeName)
	var dbInfos []db.Settings
	for _, v := range sqlServerArray.([]interface{}) {
		confMap := v.(map[string]interface{})
		dbInfo := db.Settings{
			DriverName:      confMap["driver-name"].(string),
			DbName:          confMap["db-name"].(string),
			DataSourceName:  confMap["data-source-name"].(string),
			MaxOpenConn:     confMap["max-open-conn"].(int),
			MaxIdleConn:     confMap["max-idle-conn"].(int),
			MaxConnLifeTime: confMap["max-conn-life-time"].(int),
		}
		dbInfos = append(dbInfos, dbInfo)
	}
	return dbInfos
}

func GetRedisSettings() redis.Settings {
	settings := redis.Settings{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"), // no password set
		Db:       viper.GetInt("redis.db"),          // use default DB
		PoolSize: viper.GetInt("redis.pool-size"),   // 连接池大小
		Timeout:  viper.GetInt("redis.timeout"),     // 超时
	}
	return settings
}
