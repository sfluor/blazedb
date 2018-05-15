package server

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port          uint   `toml:"port"`
	MaxQueueSize  uint   `toml:"max_queue_size"`
	LogFile       string `toml:"log_file"`
	SaveDirectory string `toml:"save_directory"`
}

func LoadConfig(filePath string) *Config {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read config file: %s", err)
		os.Exit(1)
	}

	conf := &Config{}
	if _, err := toml.Decode(string(data), conf); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't decode config file: %s", err)
		os.Exit(1)
	}

	if conf.Port == 0 {
		conf.Port = 9876
	}

	if conf.MaxQueueSize == 0 {
		conf.MaxQueueSize = 100
	}

	if conf.LogFile == "" {
		conf.LogFile = "/tmp/blazedb.log"
	}

	if conf.SaveDirectory == "" {
		conf.SaveDirectory = "/tmp"
	}

	return conf
}
