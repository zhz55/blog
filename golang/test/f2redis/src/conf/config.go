package conf

import (
    "github.com/Unknown/goconfig"
    "log"
)

type Config struct {
    configParser *goconfig.ConfigFile
    redisHost string
    redisPort uint8
    fileName string
    
}

func NewConfig()
