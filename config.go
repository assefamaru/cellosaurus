package cellosaurus

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// Context is an API configuration struct.
type Context struct {
	Mode    string
	Port    string
	Version string
}

// API server environment modes.
const (
	DebugMode   string = "debug"   // for development
	ReleaseMode string = "release" // for production
	TestMode    string = "test"    // for testing
)

// Mode, port, version global settings.
var (
	apiMode    string
	apiPort    string
	apiVersion string
)

// SetMode sets api mode.
func SetMode(mode string) {
	switch mode {
	case DebugMode:
		apiMode = DebugMode
	case ReleaseMode:
		apiMode = ReleaseMode
	case TestMode:
		apiMode = TestMode
	default:
		panic(fmt.Errorf("API mode '%s' not recognized", mode))
	}
	// Set gin mode.
	gin.SetMode(mode)
}

// SetPort sets api port.
func SetPort(port string) {
	apiPort = port
}

// SetVersion sets api version.
func SetVersion(version string) {
	apiVersion = version
}

// GetEnvMode gets api mode from environment variable MODE.
func GetEnvMode() string {
	m := os.Getenv("MODE")
	if m == "" {
		panic("MODE environment variable does not exist.")
	}
	return m
}

// GetEnvPort gets api port from environment variable PORT.
func GetEnvPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		panic("PORT environment variable does not exist.")
	}
	return p
}

// GetEnvVersion gets api version from environment variable VERSION.
func GetEnvVersion() string {
	v := os.Getenv("VERSION")
	if v == "" {
		panic("VERSION environment variable does not exist.")
	}
	return v
}

// Mode returns current api mode.
func Mode() string {
	return apiMode
}

// Port returns current api port.
func Port() string {
	return apiPort
}

// Version returns current api version.
func Version() string {
	return apiVersion
}
