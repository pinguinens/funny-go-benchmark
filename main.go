package main

import (
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/runner"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

func main() {
	appCfg := config.New()
	logger := log.Logger{}

	rn := runner.New(appCfg, &logger)
	rn.Init()
	rn.Run()
}
