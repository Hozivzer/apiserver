package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	//初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//监控配置文件变化并热加载
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) //如果指定了配文件地址，则解析指定文件
	} else {
		viper.AddConfigPath("conf") //如果未指定哦欸指纹机，则接卸默认文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     //设置配置文件格式 yaml
	viper.AutomaticEnv()            //读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") //读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

//监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed :%s", e.Name)
	})
}
