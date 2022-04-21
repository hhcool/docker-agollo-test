package config

import (
	"fmt"
	"github.com/hhcool/gtls/log"
	"github.com/hhcool/gtls/utils"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/shima-park/agollo"
	remote "github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
)

var Cfg Config

func InitConfig() {
	if endpoint := os.Getenv("APOLLO_URL"); endpoint != "" {
		initConfigFromApollo(endpoint)
	} else {
		initConfigFromFile()
	}
}
func initConfigFromFile() {
	f := DefaultConfigFile
	if gin.Mode() == gin.DebugMode {
		f = DefaultDevConfigFile
	}
	log.Infof("%-10s[%s]", "加载配置文件", f)
	v := viper.New()
	v.SetConfigFile(f)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件加载错误: %v", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("%-10s[%s]", "配置文件变化", e.Name)
		if err := v.Unmarshal(&Cfg); err != nil {
			panic(fmt.Errorf("配置文件解析错误: %v", err))
		}
	})
	if err := v.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("配置文件解析错误: %v", err))
	}
}
func initConfigFromApollo(endpoint string) {
	log.Infof("%-10s[%s]", "使用远程配置", endpoint)
	appid := utils.FirstReal(os.Getenv("APOLLO_PROJECT"), "edp").(string)
	log.Infof("%-10s[%s]", "远程配置项目", appid)
	cluster := utils.FirstReal(os.Getenv("APOLLO_CLUSTER"), "bty").(string)
	log.Infof("%-10s[%s]", "远程配置集群", cluster)

	remote.SetAppID(appid)
	remote.SetAgolloOptions(
		agollo.Cluster(cluster),
		agollo.AutoFetchOnCacheMiss(),
		agollo.FailTolerantOnBackupExists(),
	)

	v := viper.New()
	v.SetConfigType("yml")
	log.Infof("%-10s[%s]", "远程配置空间", DefaultApolloNamespace)
	if err := v.AddRemoteProvider("apollo", endpoint, DefaultApolloNamespace); err != nil {
		panic(err)
	}
	log.Infof("%-10s[%s]", "开始拉取配置", "请等待...")
	if err := v.ReadRemoteConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&Cfg); err != nil {
		panic(err)
	}
	log.Infof("%-10s[%s]", "配置内容概览", utils.StructToString(Cfg))
}
