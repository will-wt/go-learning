package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Datasource Datasource `yaml:"datasource" mapstructure:"datasource"`

	// 其它配置项
}

// LoadConfig 加载应用的 yaml 配置项
func LoadConfig() *AppConfig {
	configFile := "3-gorm/config/app.yml"
	// 初始化实例
	config := &AppConfig{}

	// 读取配置文件 os 方式，然后使用 yaml 提取配置项
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("配置文件读取异常：%s \n", err))
	}

	// 解析 yaml 配置项到已定义的结构体
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Errorf("配置文件读取异常：%s \n", err)
	}

	// 读取配置文件 viper 方式，使用 mapstructure 标签映射属性
	//viper.SetConfigFile(configFile)
	//err := viper.ReadInConfig()
	//if err != nil {
	//	panic(fmt.Errorf("配置文件读取异常：%s \n", err))
	//}
	//
	//// 解析配置项到已定义的结构体
	//err = viper.Unmarshal(&Config)
	//if err != nil {
	//	panic(fmt.Errorf("配置文件读取异常：%s \n", err))
	//}

	fmt.Printf("AppConfig: %+v \n", config)
	return config
}
