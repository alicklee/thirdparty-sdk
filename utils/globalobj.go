package utils

/**
Framework的全局框架配置
*/

type GlobalObj struct {
	//当前服务器监听的IP
	Host string
	//当前服务器监听的TCP端口
	TcpPort int
	//服务器的名字
	Name string
	//最大客户端连接数
	MaxConn int
	//最大包体长度
	MaxPackageSize uint16
	//版本号
	Version string
	//工作池的最大连接数
	WorkerPoolSize uint32
	//任务队列最大等待数量
	TaskQueueMaxLen uint32
	//token种子
	TokenSeed string
	//MongoDBUrl
	MongoDBUrl string
	//MongoDB database name
	MongoDBName string
	//日志路径
	LogPath string
	//日志等级
	LogLevel string
	//日志输出
	LogOutput string
}

//定义一个全局的配置
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	GlobalObject = g
}

//func (g *GlobalObj) Reload() {
//	data, err := conf.Asset("conf/global.json")
//	if err != nil {
//		panic(err)
//	}
//	err = json.Unmarshal(data, &GlobalObject)
//	if err != nil {
//		panic(err)
//	}
//}

func InitGlobalObject() {
	//如果配置文件没有加载的默认值
	GlobalObject = &GlobalObj{
		Host:            "127.0.0.1",
		TcpPort:         1125,
		Name:            "User Center",
		MaxConn:         1000,
		MaxPackageSize:  4096,
		Version:         "0.0.1",
		WorkerPoolSize:  10,
		TaskQueueMaxLen: 1024,
		TokenSeed:       "Cl0udC@de",
		LogPath:         GetenvDefault("LOG_FILE", "/tmp"),
		LogLevel:        GetenvDefault("LOG_LEVEL", "debug"),
		LogOutput:       GetenvDefault("LOG_OUTPUT", "std"),
		MongoDBName:     GetenvDefault("MONGO_DB_NAME", "demo"),
		MongoDBUrl:      GetenvDefault("MONGO_DB_URL", ""),
	}
	//GlobalObject.Reload()
}
