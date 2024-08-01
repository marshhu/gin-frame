package config

import (
	"encoding/json"
	"github.com/marshhu/gin-frame/common"
	"github.com/marshhu/gin-frame/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type EnvConf struct {
	UseRemoteConf  string `json:"USE_REMOTE_CONF"`
	CfgFileName    string `json:"CFG_FILE_NAME"`
	CfgGroup       string `json:"CFG_GROUP"`
	CfgFileType    string `json:"CFG_FILE_TYPE"`
	CfgPath        string `json:"CFG_PATH"`
	CfgDataID      string `json:"CFG_DATA_ID"`
	CfgNameSpaceID string `json:"CFG_NAME_SPACE_ID"`
	CfgServer      string `json:"CFG_SERVER"`
}

func GetEnvConf(envType common.EnvType) (*EnvConf, error) {
	log.Printf("envType=%s", envType)
	if envType == common.LocalFileEnvType {
		envFile := filepath.Join(utils.RootDir(), "env.json")
		if utils.CheckNotExist(envFile) { // 如果文件不存在
			envFile = filepath.Join(utils.ExecutablePath(), "env.json")
		}
		content, err := utils.ReadFile(envFile)
		if err != nil {
			return nil, err
		}
		evnConf := &EnvConf{}
		if err = json.Unmarshal(content, evnConf); err != nil {
			return nil, err
		}
		log.Printf("envInfo=%v", *evnConf)
		return evnConf, nil
	} else {
		evnConf := &EnvConf{
			UseRemoteConf:  strings.TrimSpace(os.Getenv(common.UseRemoteConf)),
			CfgFileName:    strings.TrimSpace(os.Getenv(common.CfgFileName)),
			CfgGroup:       strings.TrimSpace(os.Getenv(common.CfgGroup)),
			CfgFileType:    strings.TrimSpace(os.Getenv(common.CfgFileType)),
			CfgPath:        strings.TrimSpace(os.Getenv(common.CfgPath)),
			CfgDataID:      strings.TrimSpace(os.Getenv(common.CfgDataID)),
			CfgNameSpaceID: strings.TrimSpace(os.Getenv(common.CfgNameSpaceID)),
			CfgServer:      strings.TrimSpace(os.Getenv(common.CfgServer)),
		}
		log.Printf("envInfo=%v", *evnConf)
		return evnConf, nil
	}
}
