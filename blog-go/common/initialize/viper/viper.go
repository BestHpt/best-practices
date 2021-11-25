/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package viper

import (
	"best-practics/common/config"
	"best-practics/common/consts"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Init(path ...string) *viper.Viper {
	var configFile string
	if len(path) == 0 {
		flag.StringVar(&configFile, "c", "", "choose configFile file.")
		flag.Parse()
		if configFile == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(consts.ConfigEnv); configEnv == "" {
				configFile = consts.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", consts.ConfigFile)
			} else {
				configFile = configEnv
				fmt.Printf("您正在使用GlobalConfig环境变量,config的路径为%v\n", configFile)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", configFile)
		}
	} else {
		configFile = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", configFile)
	}
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configFile file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("configFile file changed:", e.Name)
		if err := v.Unmarshal(config.ConfigCenter); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(config.ConfigCenter); err != nil {
		fmt.Println(err)
	}
	return v
}
