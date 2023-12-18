package main

import (
	"github.com/marshhu/gin-frame/app"
	"github.com/marshhu/gin-frame/common"
	"github.com/marshhu/gin-frame/router"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := app.InitConf(common.EnvironmentEnvType); err != nil {
		log.Println(err)
		panic(err)
	}
	//err := app.InitModules([]string{common.SqlServerModule, common.RedisModule})
	err := app.InitModules([]string{common.MysqlModule, common.RedisModule})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	config := app.ServerConfig{
		HostPort:     viper.GetInt("host.port"),
		ReadTimeout:  viper.GetInt("host.read-timeout"),
		WriteTimeout: viper.GetInt("host.write-timeout"),
		RunMode:      viper.GetString("host.run-mode"),
	}
	application := app.NewApp(config)
	application.RegisterRouter(router.Default)
	application.Run(func() {
		log.Println("Server exited")
	})
}
