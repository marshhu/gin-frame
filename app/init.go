package app

import (
	"fmt"
	"github.com/marshhu/gin-frame/common"
	"github.com/marshhu/gin-frame/config"
	"github.com/marshhu/gin-frame/db"
	"github.com/marshhu/gin-frame/log"
	"github.com/marshhu/gin-frame/redis"
	"github.com/marshhu/gin-frame/utils"
	"github.com/spf13/viper"
	sysLog "log"
	"strings"
)

// InitModules 模块初始化
func InitModules(modules []string) error {
	// 初始化日志
	logSetting := GetLogSettings()
	log.Init(&logSetting)
	// 初始化sql-server
	if utils.StringInArray(common.SqlServerModule, modules) {
		dbSettings, err := GetDbSettings(common.SqlServerModule)
		if err != nil {
			return err
		}
		if err = db.InitSqlServer(dbSettings); err != nil {
			return err
		}
	}
	// 初始化mysql
	if utils.StringInArray(common.MysqlModule, modules) {
		dbSettings, err := GetDbSettings(common.MysqlModule)
		if err != nil {
			return err
		}
		if err := db.InitMysql(dbSettings); err != nil {
			return err
		}
	}
	// 初始化redis
	if utils.StringInArray(common.RedisModule, modules) {
		redisSettings := GetRedisSettings()
		if err := redis.InitClient(redisSettings); err != nil {
			return err
		}
	}
	return nil
}

func InitConf(envType common.EnvType) error {
	envInfo, err := config.GetEnvConf(envType)
	if err != nil {
		return err
	}
	if envInfo == nil || len(envInfo.UseRemoteConf) == 0 {
		return fmt.Errorf("请检查env变量配置是否正确：envType = %s", envType)
	}

	if strings.ToLower(envInfo.UseRemoteConf) == "true" {
		if err = config.InitRemote(envInfo); err != nil {
			sysLog.Printf("加载远程配置失败：%v", err)
			return err
		}
		sysLog.Println("当前使用的配置是：远程配置")
	} else {
		err = config.InitLocal(envInfo)
		if err != nil {
			sysLog.Printf("加载本地配置失败：%v", err)
			return err
		}
		sysLog.Println("当前使用的配置是：本地配置")
	}
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

func GetDbSettings(confNodeName string) ([]db.Settings, error) {
	sqlServerArray := viper.Get(confNodeName)
	if sqlServerArray == nil {
		return nil, fmt.Errorf("读取%s配置失败", confNodeName)
	}
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
	return dbInfos, nil
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
