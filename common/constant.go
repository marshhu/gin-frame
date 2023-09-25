package common

type EnvType string

const (
	EnvironmentEnvType EnvType = "environment"
	LocalFileEnvType   EnvType = "local-file"
)
const (
	UseRemoteConf  = "USE_REMOTE_CONF"   // 是否启用远程配置
	CfgServer      = "CFG_SERVER"        // 远程配置服务地址
	CfgNameSpaceID = "CFG_NAME_SPACE_ID" // 命名空间ID
	CfgDataID      = "CFG_DATA_ID"       //  数据集ID
	CfgGroup       = "CFG_GROUP"         // 分组
	CfgFileName    = "CFG_FILE_NAME"     // 本地配置文件名
	CfgFileType    = "CFG_FILE_TYPE"     // 本地配置文件类型
	CfgPath        = "CFG_PATH"          // 本地配置文件路径
)

const (
	SqlServerModule = "sql-server"
	MysqlModule     = "mysql"
	RedisModule     = "redis"
)
