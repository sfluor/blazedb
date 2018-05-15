package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// Config struct
type Config struct {
	Port         uint          `toml:"port"`
	MaxQueueSize uint          `toml:"max_queue_size"`
	LogFile      string        `toml:"log_file"`
	SaveFile     string        `toml:"save_file"`
	Debug        uint          `toml:"debug"`
	SavePeriod   time.Duration `toml:"save_period"`
}

// LoadConfig loads the configuration from the given file path and returns it
// It takes care of setting some attributes to their default values if they are not set
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

	if conf.SaveFile == "" {
		conf.SaveFile = "/tmp/db.blaze"
	}

	if conf.SavePeriod == 0 {
		conf.SavePeriod = 1 * time.Minute
	}

	return conf
}
