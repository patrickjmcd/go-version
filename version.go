package version

import "fmt"

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"
)

func GetVersion() string {
	return fmt.Sprintf("commit: %s, build time: %s, release: %s",
		Commit, BuildTime, Release)
}
