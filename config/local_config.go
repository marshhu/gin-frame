package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
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

func InitLocal() {
	config := Setting{
		CfgFile: getConfigName(),
		CfgDirs: getConfigDirs(),
		CfgType: getConfigType(),
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
		panic(fmt.Errorf("加载配置文件失败: %s \n", err))
	}
	log.Println("使用的配置文件位置：" + viper.ConfigFileUsed())
}

func getConfigName() string {
	cfgName := strings.TrimSpace(os.Getenv(CfgFileName))
	if len(cfgName) == 0 {
		cfgName = defaultCFGName
	}

	log.Printf("CFG_FILE=%s\r\n", cfgName)

	// 根据环境变量加载不同的配置文件
	env := strings.TrimSpace(os.Getenv(CfgGroup))

	if len(env) == 0 {
		env = defaultCFGEnv
	}
	log.Printf("%s=%s\r	\n", CfgGroup, env)

	cfgFileName := cfgName + "." + env

	return cfgFileName
}

func getConfigDirs() []string {
	// 添加多个配置文件路径，以先找到的为准
	rootDir, _ := os.Getwd()
	configDir := filepath.FromSlash(path.Join(rootDir, "conf"))
	envDir := os.Getenv(CfgPath)
	log.Printf("%s=%s\r\n", CfgPath, envDir)

	etcDirs := []string{envDir, configDir, rootDir}
	return etcDirs
}

func getConfigType() string {
	cfgFileType := strings.TrimSpace(os.Getenv(CfgFileType))
	if len(cfgFileType) == 0 {
		cfgFileType = defaultCFGType
	}
	return cfgFileType
}
