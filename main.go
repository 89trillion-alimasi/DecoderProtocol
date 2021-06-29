package main

import (
	"DecoderProtocol/config"
	"DecoderProtocol/envs"
	"DecoderProtocol/router"
	"fmt"
	"github.com/sirupsen/logrus"
)

// configPath app.ini 配置文件位置, pflag 会将 "--config" 参数设置到这个变量中
var configPath string

// inputPath outputPath 数据文件的输入和输出位置, 同样 pflag 读取 "--input" 和 "--output" 设置
var inputPath, outputPath string

// help 参数控制是否打印帮助信息
var help bool

// main 主执行方法
func main() {

	// 解析命令行参数, 并读取 ini 中的端口号
	// 如果解析失败会立即退出
	port := envs.Parse()

	// 解析数据文件
	config.ParseJson(envs.Conf.InputPath, envs.Conf.OutputPath)

	// 初始化 gin router
	r := router.InitRouter()

	// 启动 gin router
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Error(err)
	}
}
