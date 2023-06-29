package settings

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	Port        int    `mapstructure:"port"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filepath string) (err error) {
	// 设置默认值
	//viper.SetDefault("fileDir", "./")
	// 读取配置文件
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 配合远程配置中心使用，etcd或者consul。告诉viper使用什么格式解析传过来的数据
	//viper.SetConfigFile("config.yaml") // 指定配置文件路径
	//viper.SetConfigFile("config.json") // 指定配置文件路径
	viper.SetConfigFile(filepath) // 指定配置文件路径
	////viper.AddConfigPath("/home/sherwood/go/src/viper_demo") // 查找配置文件所在的路径
	////viper.AddConfigPath("$HOME/.appName")                   // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")   // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		//zap.L().Error("viper.ReadInConfig() failed,err:%v\n", zap.Error(err))
		return err
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 将读取到的信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
		zap.L().Info("config file changed")
		if err := viper.Unmarshal(Conf); err != nil {
			return
		}
	})
	return err
}
