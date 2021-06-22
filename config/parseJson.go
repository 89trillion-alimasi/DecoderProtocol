package config

import (
	"DecoderProtocol/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

// ArmyData 是一个 map，用于从输入的 json 中反序列化存储士兵信息
type ArmyData map[string]model.ArmyInfo

// AArmyData 存储解析输入的 json 后，所有士兵有用的信息
var AArmyData = make(ArmyData)

// ParseJson 对输入的 json 文件(inputPath)进行解析
// 并将解析结果(保留所有有用信息)写入到输出文件(outputPath)中
func ParseJson(inputPath, outputPath string) {

	// 读取配置数据文件
	bs, err := ioutil.ReadFile(inputPath)
	if err != nil {
		logrus.Fatalf("无法读取输入的 json 数据(%s): %v", inputPath, err)
	}

	// 反序列化 json，将其转换并存储到 AArmyData 中
	err = jsoniter.Unmarshal(bs, &AArmyData)
	if err != nil {
		logrus.Fatalf("无法解析输入的 json 数据(%s): %v", inputPath, err)
	}
	// 重新序列化 AArmyData，得到 "保留数据" 的 json
	bs, err = jsoniter.MarshalIndent(AArmyData, "", "    ")
	if err != nil {
		logrus.Fatalf("重新序列化 json 数据失败(%s): %v", outputPath, err)
	}

	// 将 "保留数据" json 写入到输出文件中
	err = ioutil.WriteFile(outputPath, bs, 0644)
	if err != nil {
		logrus.Fatalf("写入 json 数据失败(%s): %v", outputPath, err)
	}

}
