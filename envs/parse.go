package envs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	// app.ini 配置文件位置, pflag 会将 "--config" 参数设置到这个变量中
	ConfigPath string
	// 数据文件的输入位置, pflag 读取 "--input" 参数设置到这个变量中
	InputPath string
	// 数据文件的输出位置, pflag 读取 "--output" 参数设置到这个变量中
	OutputPath string
	// 是否打印帮助信息, pflag 读取 "--help" 参数设置到这个变量中
	Help bool
}

// Conf 存储全局命令行参数
var Conf Config

// Parse 执行后会解析命令行参数，并将值设置到 Conf 变量中
func Parse() int {

	//设置 logrus 日志格式
	logrus.SetFormatter(&logrus.TextFormatter{
		// 打印完整时间戳
		FullTimestamp: true,
		// 时间戳格式 yyyy-MM-dd hh:mm:ss
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// pflag 命令行参数定义
	pflag.StringVarP(&Conf.ConfigPath, "config", "c", "./app.ini", "配置文件位置")
	pflag.StringVarP(&Conf.InputPath, "input", "i", "./config.army.model.json", "输入的 json 数据文件")
	pflag.StringVarP(&Conf.OutputPath, "output", "o", "./config.gen.json", "输出的 json 数据文件")
	pflag.BoolVarP(&Conf.Help, "Help", "h", false, "打印帮助信息")

	// 调用 pflag 进行解析
	pflag.Parse()

	// 打印帮助信息
	if Conf.Help {
		fmt.Printf("%s --config CONFIG_INI_FILE --input INPUT_JSON_FILE --output OUTPUT_JSON_FILE\n", os.Args[0])
		// 打印帮助信息后不需要启动服务，直接推出即可
		os.Exit(0)
	}

	// 读取配置文件
	conf, err := ini.Load(Conf.ConfigPath)
	if err != nil {
		logrus.Fatalf("无法读取配置文件 [%s]: %v", Conf.ConfigPath, err)
	}

	// 获取 HttpPort
	key, err := conf.Section("server").GetKey("HttpPort")
	if err != nil || key == nil {
		logrus.Fatalf("无法读取配置文件 [%s]: key HttpPort 未找到", Conf.ConfigPath)
	}
	// HttpPort 端口号必须为 int
	port, err := key.Int()
	if err != nil {
		logrus.Fatalf("无法读取配置文件 [%s]: %v", Conf.ConfigPath, err)
	}

	return port
}
