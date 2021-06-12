/*
	提供服务初始化，包括数据库，Redis，Mq，日志文件等
*/
package initialize

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	MsqlDb       *gorm.DB             //数据库客户端
	//IrisLog         *golog.Logger     //全局log变量
	Config          *viper.Viper      //全局配置
	GrpcConn        *grpc.ClientConn          //grpc对象 用于grpc调用
	Client 			client.Client
)

//	提供系统初始化，全局变量
func Init(config *viper.Viper) {

	Config = config
	var err error
	//mysql配置
	MsqlDb, err = gorm.Open("mysql", config.GetString("Mysql.user")+":"+config.GetString("Mysql.password")+"@tcp("+config.GetString("Mysql.host")+":"+config.GetString("Mysql.Port")+")/"+config.GetString("Mysql.database")+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	MsqlDb.DB().SetMaxIdleConns(10)
	MsqlDb.DB().SetMaxOpenConns(100)
	// 激活链接
	if err = MsqlDb.DB().Ping(); err != nil {
		panic(err)
	}

	////	系统日志配置
	//App = iris.New()
	//IrisLog = App.Logger()
	//App.Use(logger.New(returnLogConfig()))
	////	终端和日志文件同时输出
	//IrisLog.SetOutput(io.MultiWriter(newLogFile(), os.Stdout))
	//IrisLog.SetPrefix("[SERVER_IP:" + libs.ReturnCurrentIp() + "]" + " [SERVER_NAME:" + config.GetString("Nacos.ServiceName") + "] ")
	//IrisLog.Handle(func(value *golog.Log) (handled bool) {
	//	caller := findCaller(6) //中间件，日志中记录打印日志的文件及函数名称
	//	value.Message = fmt.Sprintf("[%s] ", caller) + value.Message
	//
	//	log := "[SERVER_IP:" + libs.ReturnCurrentIp() + "]" + " [SERVER_NAME:" + config.GetString("Nacos.ServiceName") + "] [" + time.Now().Format("2006-01-02 15:04:05") + "] " + value.Message
	//
	//	//判断有没有ServeHTTP
	//	if strings.Contains(log, "ServeHTTP") {
	//		//如果包含就退出
	//		return false
	//	}
	//	return false
	//})

	//Client
	Client = micro.NewService().Client()

}

//	查找调用包名，文件名，函数名 方便日志查找
func findCaller(skip int) string {
	file := ""
	line := 0
	var pc uintptr
	// 遍历调用栈的最大索引为第11层.
	for i := 0; i < 11; i++ {
		file, line, pc = getCaller(skip + i)
		// 过滤掉所有logrus包，即可得到生成代码信息
		if !strings.HasPrefix(file, "golog") {
			break
		}
	}

	fullFnName := runtime.FuncForPC(pc)

	fnName := ""
	if fullFnName != nil {
		fnNameStr := fullFnName.Name()
		// 取得函数名
		parts := strings.Split(fnNameStr, ".")
		fnName = parts[len(parts)-1]
	}

	return fmt.Sprintf("%s:%d:%s()", file, line, fnName)
}

//	查找调用包名，文件名，函数名 方便日志查找
func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}
	n := 0

	// 获取包名
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, pc
}

//	设置日志配置
//func returnLogConfig() logger.Config {
//	return logger.Config{
//		// Status displays status code
//		Status: true,
//		// IP displays request's remote address
//		IP: true,
//		// Method displays the http method
//		Method: true,
//		// Path displays the request path
//		Path: true,
//		// Query appends the url query to the Path.
//		Query: true,
//		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
//		// will be added to the logs.
//		MessageContextKeys: []string{"logger_message"},
//		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
//		MessageHeaderKeys: []string{"User-Agent"},
//	}
//}

//	创建日志文件目录和文件
func newLogFile() *os.File {
	path := "./logs/"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic("创建日志logs文件失败")
		}
	}
	filename := path + time.Now().Format("20060102") + ".log"
	// Open the file, this will append to the today's file if server restarted.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

//注册nacos
//func InitRegisterServer() {
//
//	// 可以没有，采用默认值
//	clientConfig := constant.ClientConfig{
//		TimeoutMs:      10 * 1000,
//		ListenInterval: 30 * 1000,
//		BeatInterval:   5 * 1000,
//	}
//
//	//获得配置url
//	url := Config.GetStringSlice("Nacos.url")
//	port := Config.GetInt("Nacos.port")
//	var serverConfigs []constant.ServerConfig
//	for _, val := range url {
//		if val != "" {
//			singleConfigs := constant.ServerConfig{
//				IpAddr:      val,          //nacos服务的ip地址
//				ContextPath: "/nacos",     //nacos服务的上下文路径，默认是“/nacos”
//				Port:        uint64(port), //nacos服务端口
//			}
//			serverConfigs = append(serverConfigs, singleConfigs)
//		}
//	}
//
//	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
//		"clientConfig":  clientConfig,
//		"serverConfigs": serverConfigs,
//	})
//
//	if err != nil {
//		IrisLog.Infof("[Nacos注册失败-%s]", err.Error())
//		return
//	}
//
//	//转换 port参数
//	portString := *configs.EnvPort
//	nacosPort, _ := strconv.Atoi(portString[1:len(portString)])
//	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
//		Ip:          libs.ReturnCurrentIp(),
//		Port:        uint64(nacosPort),
//		ServiceName: Config.GetString("Nacos.ServiceName"),
//		Weight:      1,
//		Enable:      true,
//		Healthy:     true,
//		Ephemeral:   true,
//		Metadata: map[string]string{
//			"content":                   "/",
//			"preserved.register.source": "SPRING_CLOUD",
//		},
//	})
//
//	if err != nil {
//		IrisLog.Info(err)
//		return
//	}
//
//	if success != true {
//		IrisLog.Infof("[Nacos注册失败-%s]", err.Error())
//		return
//	}
//
//	//赋值nacos服务
//	NacosClient = namingClient
//
//	IrisLog.Info("[Nacos注册成功]")
//}
