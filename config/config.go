package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	SERVER_IP string
	PORT      string
}

func GetConfig(params ...string) (config Configuration) {
	config = Configuration{}

	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	fileName := fmt.Sprintf("./config/%s_config.yaml", env)

	gonfig.GetConf(fileName, &config)

	return
}
