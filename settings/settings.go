package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	// 设置默认值
	//viper.SetDefault("fileDir", "./")
	// 读取配置文件
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	////viper.AddConfigPath("/home/sherwood/go/src/viper_demo") // 查找配置文件所在的路径
	////viper.AddConfigPath("$HOME/.appName")                   // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")   // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return err
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
	})
	return err
}
