package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/marshhu/gin-frame/utils"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

func InitRemote(envInfo *EnvConf) error {
	serverUrl, err := url.Parse(envInfo.CfgServer)
	if err != nil {
		return err
	}
	if len(envInfo.CfgNameSpaceID) == 0 {
		return errors.New("未配置命名空间ID")
	}

	if len(envInfo.CfgDataID) == 0 {
		return errors.New("未配置DataID")
	}
	if len(envInfo.CfgGroup) == 0 {
		return errors.New("未配置group")
	}

	if len(envInfo.CfgFileType) == 0 {
		return errors.New("未配置file type")
	}
	serverUrlArray := strings.Split(serverUrl.Host, ":")
	if len(serverUrlArray) != 2 {
		return errors.New("CFG_SERVER配置格式错误")
	}
	port, err := strconv.ParseInt(serverUrlArray[1], 10, 64)
	if err != nil {
		return err
	}
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(serverUrlArray[0], uint64(port), constant.WithContextPath("/nacos")),
	}
	logDir := filepath.Join(utils.RootDir(), "nacos/log")
	catchDir := filepath.Join(utils.RootDir(), "nacos/cache")
	if err = utils.MkDir(logDir); err != nil {
		return err
	}
	if err = utils.MkDir(catchDir); err != nil {
		return err
	}
	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(envInfo.CfgNameSpaceID),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(logDir),
		constant.WithCacheDir(catchDir),
		constant.WithLogLevel("debug"),
		//common.WithUsername("nacos"),
		//common.WithPassword("Hzz880719"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return err
	}
	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: envInfo.CfgDataID,
		Group:  envInfo.CfgGroup,
	})
	if err != nil {
		return err
	}
	fmt.Println("GetConfig,config :" + content)
	viper.SetConfigType(envInfo.CfgFileType)
	viper.ReadConfig(bytes.NewBuffer([]byte(content)))
	// 监听变化
	err = client.ListenConfig(vo.ConfigParam{
		DataId: envInfo.CfgDataID,
		Group:  envInfo.CfgGroup,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("namespace:" + namespace + "group:" + group + ", dataId:" + dataId + ", data:" + data)
			viper.ReadConfig(bytes.NewBuffer([]byte(data)))
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
