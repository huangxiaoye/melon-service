package main

import (
	"log"
	"melon-service/databases"
	"melon-service/routers"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) //读取配置文件失败致命错误
	}
	if viper.GetString("run_mode") == "debug" {
		os.Setenv("UPPERIO_DB_DEBUG", "1")
	}
}

func main() {
	databases.Init()
	r := routers.Engine()
	r.Run()
}
