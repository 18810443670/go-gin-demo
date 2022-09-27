package Services

import (
	"github.com/go-ini/ini"
	"log"
	"path/filepath"
)

var Config *ini.File

func InitIni() *ini.File {
	absPath, _ := filepath.Abs("config.ini")
	Config, err := ini.Load(absPath)
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	return Config
}
