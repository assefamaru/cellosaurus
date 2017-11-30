package cellosaurus

import raven "github.com/getsentry/raven-go"

// Sentry DSN for internal error logging.
func init() {
	raven.SetDSN("https://36b98457994b46efb1dea6c9ffd9eb70:19a5e80e08e043aeb6ef9f60693bbcf9@sentry.io/156124")
}

// LogSentry submits internal errors to Sentry.
func LogSentry(err error) {
	raven.CaptureError(err, nil)
}
