package config

import "os"

type Configuration struct {
	Server Server
	Media  Media
}

type Server struct {
	Port string `yaml:"port" default:"8080"`
}

type Media struct {
	Dir string `yaml:"dir" default:"/tmp"`
}

func InitConfiguration() Configuration {
	return LoadEnvironmentVariables()
}

func LoadEnvironmentVariables() Configuration {

	mediaDir := os.Getenv("MEDIA_DIR")

	if mediaDir == "" {
		mediaDir = "/tmp"
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return Configuration{
		Server: Server{
			Port: port,
		},
		Media: Media{
			Dir: mediaDir,
		},
	}

}
