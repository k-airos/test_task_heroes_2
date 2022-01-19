package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"marvel/internal/app/api"
)

var (
	ConfigPath string
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&ConfigPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	//В этот момент происходит инициализация переменной configPath
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		config.Port = ":8081"
		log.Println("Can not find conf file, using default value")
	}
	server := api.New(config)
	server.Config.ConfigPath = ConfigPath
	api.Server = server
	api.ConfigureStorageField()
	api.ConfigureRouterField() //gin.Run() here

}
