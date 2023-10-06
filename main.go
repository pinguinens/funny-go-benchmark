package main

import (
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/service"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

func main() {
	appCfg := config.New()
	logger := log.Logger{}

	svc := service.New(appCfg, &logger)
	svc.Init()
	svc.Run()
}
