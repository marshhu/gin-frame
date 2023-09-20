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
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func initRemote() error {
	confServer := strings.TrimSpace(os.Getenv(CfgServer))
	serverUrl, err := url.Parse(confServer)
	log.Printf("%s=%s", CfgServer, serverUrl)
	if err != nil {
		return err
	}
	nameSpaceId := strings.TrimSpace(os.Getenv(CfgNameSpaceID))
	log.Printf("%s=%s", CfgNameSpaceID, nameSpaceId)
	if len(nameSpaceId) == 0 {
		return errors.New("未配置命名空间ID")
	}
	dataID := strings.TrimSpace(os.Getenv(CfgDataID))
	log.Printf("%s=%s", CfgDataID, dataID)
	if len(nameSpaceId) == 0 {
		return errors.New("未配置DataID")
	}
	group := strings.TrimSpace(os.Getenv(CfgGroup))
	log.Printf("%s=%s", CfgGroup, group)
	if len(nameSpaceId) == 0 {
		return errors.New("未配置group")
	}
	fileType := strings.TrimSpace(os.Getenv(CfgFileType))
	log.Printf("%s=%s", CfgFileType, fileType)
	if len(nameSpaceId) == 0 {
		fileType = "yaml"
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
		constant.WithNamespaceId(nameSpaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(logDir),
		constant.WithCacheDir(catchDir),
		constant.WithLogLevel("debug"),
		//constant.WithUsername("nacos"),
		//constant.WithPassword("Hzz880719"),
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
		DataId: dataID,
		Group:  group,
	})
	if err != nil {
		return err
	}
	fmt.Println("GetConfig,config :" + content)
	viper.SetConfigType(fileType)
	viper.ReadConfig(bytes.NewBuffer([]byte(content)))
	// 监听变化
	err = client.ListenConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
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
