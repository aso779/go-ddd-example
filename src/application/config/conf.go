package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	AppName  string     `yaml:"app_name"`
	Port     string     `yaml:"port"`
	Logs     Logs       `yaml:"logs"`
	Postgres PostgresRW `yaml:"postgres"`
}

const usageText = `
Supported commands are:
-c [filename] - start application with specified configuration
`

const requiredText = `
-c [filename] - arg not specified
`

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

var conf *Config

func NewConfig() *Config {
	if conf != nil {
		return conf
	}

	flag.Usage = usage
	c := flag.String("c", "", "configuration filename")
	flag.Parse()

	if *c == "" {
		fmt.Print(requiredText)
		os.Exit(2)
	}

	config := Config{}

	file, err := ioutil.ReadFile(*c)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(128)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Print(fmt.Sprintf("%s %s", "load config error", err.Error()))
		os.Exit(128)
	}

	conf = &config

	return conf
}
