package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"time"
)

func main() {
	NacosConfig := &NacosConfig{}
	debug := true
	configFileName := fmt.Sprintf("config-pro.yaml")
	if debug {
		configFileName = fmt.Sprintf("config-debug.yaml")
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(NacosConfig); err != nil {
		panic(err)
	}
	sc := []constant.ServerConfig{
		{
			IpAddr: NacosConfig.Host,
			Port:   NacosConfig.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         NacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConfig.DataId,
		Group:  NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = yaml.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: NacosConfig.DataId,
		Group:  NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	time.Sleep(3000 * time.Second)

}
