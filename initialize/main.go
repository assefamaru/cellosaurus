package main

import (
	"flag"

	"github.com/assefamaru/cellosaurus"
)

func main() {
	var (
		ctx cellosaurus.Context

		getModeFromEnv    bool
		getPortFromEnv    bool
		getVersionFromEnv bool
	)

	// Flags
	flag.BoolVar(&getModeFromEnv, "emode", false, "use MODE environment variable")
	flag.BoolVar(&getPortFromEnv, "eport", false, "use PORT environment variable")
	flag.BoolVar(&getVersionFromEnv, "eversion", false, "use VERSION environment variable")
	flag.StringVar(&ctx.Mode, "mode", "release", "environment mode")
	flag.StringVar(&ctx.Port, "port", "8080", "server port")
	flag.StringVar(&ctx.Version, "version", "v2", "api version")

	flag.Parse()

	if getModeFromEnv {
		ctx.Mode = cellosaurus.GetEnvMode()
	}
	if getPortFromEnv {
		ctx.Port = cellosaurus.GetEnvPort()
	}
	if getVersionFromEnv {
		ctx.Version = cellosaurus.GetEnvVersion()
	}

	cellosaurus.Init(&ctx)
}
