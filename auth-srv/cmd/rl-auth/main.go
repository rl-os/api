package main

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
	buildDate  string
)

const (
	// appName is an identifier-like name used anywhere this app needs to be identified.
	//
	// It identifies the application itself, the actual instance needs to be identified via environment
	// and other details.
	appName = "rl.auth"

	// friendlyAppName is the visible name of the application.
	friendlyAppName = "RisuLife OAuth Server"
)

func main() {

}
