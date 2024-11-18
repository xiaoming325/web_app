package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 初始化viper，读取配置文件并监视更改
func Init() (err error) {
	viper.SetConfigName("config")  // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")    // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./conf/") // 查找配置文件所在的路径
	err = viper.ReadInConfig()     // 查找并读取配置文件
	if err != nil {                // 处理读取配置文件的错误
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	//监视配置文件的更改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了...")
		return
	})

	return
}
