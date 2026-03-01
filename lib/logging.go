package lib

// Global logging configuration
var (
	EnableMacLogging bool = false
	EnableStpLogging bool = false
)

// SetLoggingScope enables logging for the specified scope
func SetLoggingScope(scope string) {
	switch scope {
	case "mac":
		EnableMacLogging = true
		EnableStpLogging = false
	case "stp":
		EnableMacLogging = false
		EnableStpLogging = true
	case "all":
		EnableMacLogging = true
		EnableStpLogging = true
	case "none":
		EnableMacLogging = false
		EnableStpLogging = false
	}
}
