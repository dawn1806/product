package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {

	configSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(host + ":" + strconv.FormatInt(port, 10)),
		// 设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix(prefix),
		// 是否移除前缀，这里设置true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
		)

	// 配置初始化
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	// 加载配置
	err = config.Load(configSource)
	return config, err
}
