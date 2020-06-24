package version

import (
	"fmt"
	"runtime"
)

// package global var, value set by linker using ldflag -X.
var (
	version string
	date    string //nolint:gochecknoglobals
	commit  string //nolint:gochecknoglobals
)

// Info - version info.
type Info struct {
	Version string
	Date    string
	Commit  string
}

// GetInfo - get version information.
func GetInfo() Info {
	return Info{
		Version: version,
		Date:    date,
		Commit:  commit,
	}
}

// String - version stringifier.
func (vi Info) String() string {
	return fmt.Sprintf("%s #%s-%s-%s [%s]",
		vi.Version,
		vi.Commit,
		runtime.GOOS,
		runtime.GOARCH,
		vi.Date,
	)
}
