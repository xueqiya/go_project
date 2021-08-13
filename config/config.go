package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Server   server
	Database database
}

type server struct {
	Mode         string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

type database struct {
	Dialect      string
	User         string
	Password     string
	Host         string
	Name         string
	Protocol     string
	Charset      string
	ParseTime    string
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

var Cfg config

// init 初始化 Cfg 全局变量。
func init() {
	_, err := toml.DecodeFile("./config.toml", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
