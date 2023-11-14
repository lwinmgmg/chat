package env

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	SocketPort int    `yaml:"socket_port"`

	Postgres   DatabaseEnv `yaml:"postgres"`
	Mongo      DatabaseEnv `yaml:"mongo"`
	GrpcServer ServerEnv   `yaml:"grpc_server"`
}

type Env struct {
	Settings Settings `yaml:"settings"`
}

func GetEnv() Env {
	path, ok := os.LookupEnv("SETTING_PATH")
	if !ok {
		path = "env.yaml"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var env Env
	if err := yaml.Unmarshal(data, &env); err != nil {
		panic(err)
	}
	return env
}
