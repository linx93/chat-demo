package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// LoadConf 加载配置
func LoadConf() *Config {
	v := viper.New()
	v.SetConfigName("dev")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs/")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("加载配置文件失败:%v", err)
	}

	get := v.Get("server.bind")
	fmt.Println(get)
	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("加载配置失败:%v", err)
	}

	//开始监听
	v.WatchConfig()

	//设置监听回调函数
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件被修改:%s \n", in.String())
		if err := viper.Unmarshal(&c); err != nil {
			log.Fatalf("配置文件被修改后报错:%v", err)
		}
	})

	return &c

}
