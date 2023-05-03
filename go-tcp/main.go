package main

import (
	"awesomeProject/go-tcp/config"
	"awesomeProject/go-tcp/lib/logger"
	"awesomeProject/go-tcp/tcp"
	"fmt"
	"os"
)

const configFile string = "redis.conf"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()

}

func main() {

	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "godis",
		Ext:        "log",
		TimeFormat: "2023-05-03",
	})
	if fileExists(configFile) {
		config.SetupConfig(configFile)
	} else {
		config.Properties = defaultProperties
	}
	err := tcp.ListenAndServeWithSignal(&tcp.Config{
		Address: fmt.Sprintf("%s:%d", config.Properties.Bind,
			config.Properties.Port),
	}, tcp.MakeHandler())
	if err != nil {
		logger.Error(err)
	}
}
