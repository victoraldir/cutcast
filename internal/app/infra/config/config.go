package config

type Configuration struct {
	Server Server
}

type Server struct {
	Port string `yaml:"port" default:"8080"`
}

func InitConfiguration() Configuration {
	return Configuration{}
}
