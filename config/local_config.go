package config

import (
	"github.com/marshhu/gin-frame/common"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Setting struct {
	CfgFile string   // 配置文件名
	CfgDirs []string // 配置路径
	CfgType string   // 配置文件格式
}

var (
	defaultCFGName = "config" // 默认配置文件名
	defaultCFGEnv  = "dev"    // 默认环境
	defaultCFGType = "yaml"   // 默认配置文件格式
)

func InitLocal(envInfo *EnvConf) error {
	config := Setting{
		CfgFile: getConfigName(envInfo),
		CfgDirs: getConfigDirs(envInfo),
		CfgType: getConfigType(envInfo),
	}
	viper.SetConfigName(config.CfgFile) // name of config file (without extension)
	viper.SetConfigType(config.CfgType) // REQUIRED if the config file does not have the extension in the name
	for _, dir := range config.CfgDirs {
		if len(dir) == 0 {
			continue
		}
		viper.AddConfigPath(dir) // path to look for the config file in
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}
	log.Println("使用的配置文件位置：" + viper.ConfigFileUsed())
	return nil
}

func getConfigName(envInfo *EnvConf) string {
	cfgName := envInfo.CfgFileName
	if len(cfgName) == 0 {
		cfgName = defaultCFGName
	}
	log.Printf("CFG_FILE=%s\r\n", cfgName)

	// 根据环境变量加载不同的配置文件
	env := envInfo.CfgGroup
	if len(env) == 0 {
		env = defaultCFGEnv
	}
	log.Printf("%s=%s\r	\n", common.CfgGroup, env)

	cfgFileName := cfgName + "." + env

	return cfgFileName
}

func getConfigDirs(envInfo *EnvConf) []string {
	// 添加多个配置文件路径，以先找到的为准
	rootDir, _ := os.Getwd()
	configDir := filepath.FromSlash(path.Join(rootDir, "conf"))
	etcDirs := []string{configDir, rootDir}
	if len(envInfo.CfgPath) > 0 {
		etcDirs = append(etcDirs, envInfo.CfgPath)
	}
	return etcDirs
}

func getConfigType(envInfo *EnvConf) string {
	cfgFileType := envInfo.CfgFileType
	if len(cfgFileType) == 0 {
		cfgFileType = defaultCFGType
	}
	return cfgFileType
}
