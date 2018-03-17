package version

import (
	"fmt"
	"runtime"
)

// VersionMajor holds the release major number.
const VersionMajor = 1

// VersionMinor holds the release minor number.
const VersionMinor = 0

// VersionPatch holds the release patch number.
const VersionPatch = 0

// Version holds the combination of major, minor and patch release numbers
var Version = fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)

// BuildHash is filled with Git revision being used to build the program at linking time.
var BuildHash = ""

// BuildNumber is filled with Git revision being used to build the program at linking time.
var BuildNumber = ""

// BuildDate in ISO8901 format, is filled when building the program at linking time.
var BuildDate = "1970-01-01T00:00:00Z"

// BuildPlatform returns the OS platform that this build was compiled for.
var BuildPlatform = fmt.Sprintf("%s/%s", runtime.GOOS, rutime.GOARCH)

// GoVersion returns the version of Go compiler used during build.
var GoVersion = runtime.Version()
