package config

import (
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/prometheus/common/model"
	"gopkg.in/yaml.v3"
)

const filePath = "config/config.yaml"

//Config all config in this
type Config struct {
	Listen     string            `yaml:"listen"`
	LabelOrder []model.LabelName `yaml:"label_order"`
	DataPath   string            `yaml:"data_path"`
}

func (c *Config) LoadEnvVars() {
	if v, ok := syscall.Getenv("CONFIG_LISTEN"); ok {
		c.Listen = v
	}
}

var conf *Config

func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func loadConfig(path string) *Config {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	conf := new(Config)
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	conf.LoadEnvVars()
	return conf
}

//Get  get a congfig instance
func Get() *Config {
	if conf == nil {
		Set(filePath)
	}
	return conf
}

func Set(filePath string) {

	if !PathIsExist(filePath) {
		log.Fatalln("The configuration file does not exist")
	}

	conf = loadConfig(filePath)

}
