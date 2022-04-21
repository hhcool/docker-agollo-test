package main

import (
	"github.com/hhcool/docker-agollo/config"
	"os"
)

func main() {
	_ = os.Setenv("APOLLO_URL", "https://apollo-dev.xxx.com")
	_ = os.Setenv("APOLLO_PROJECT", "loc")
	_ = os.Setenv("APOLLO_CLUSTER", "default")

	config.InitConfig()
}
