package gomicroboilerplate

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

// Populated during build, don't touch!
var (
	Version   = "v0.1.0"
	GitRev    = "undefined"
	GitBranch = "undefined"
	BuildDate = "Fri, 17 Jun 1988 01:58:00 +0200"
)

// PrintVersion prints version info into the provided io.Writer.
func PrintVersion() {
	logrus.WithFields(logrus.Fields{
		"version":    Version,
		"git_rev":    GitRev,
		"git_branch": GitBranch,
		"build_date": BuildDate,
		"go_version": runtime.Version(),
		"os":         runtime.GOOS,
		"arch":       runtime.GOARCH,
	}).Info("version info")
}
