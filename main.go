package main

import (
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/runner"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/memory"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

func main() {
	appCfg := config.New()
	logger := log.Logger{}

	test := memory.Tester{}

	rn := runner.New(appCfg, &logger)
	rn.Run(&test)
}
