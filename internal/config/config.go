package config

import (
	"flag"
	"runtime"
)

type App struct {
	Buffer Buffer
	System System
}

type Buffer struct {
	Multiplier int
}

type System struct {
	CoreCount int
}

func New() *App {
	cfg := App{}

	flag.IntVar(&cfg.Buffer.Multiplier, "b", 1, "buffer memory size in MB")
	flag.IntVar(&cfg.System.CoreCount, "c", runtime.NumCPU(), "CPU cores")
	flag.Parse()

	return &cfg
}
