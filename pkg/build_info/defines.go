package build_info

import "runtime"

// BuildInfo represents all available build information.
type BuildInfo struct {
	Version    string `json:"version"`
	CommitHash string `json:"commit_hash"`
	BuildDate  string `json:"build_date"`
	GoVersion  string `json:"go_version"`
	Os         string `json:"os"`
	Arch       string `json:"arch"`
	Compiler   string `json:"compiler"`
}

func New(version string, commitHash string, buildDate string) BuildInfo {
	return BuildInfo{
		Version:    version,
		CommitHash: commitHash,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Os:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		Compiler:   runtime.Compiler,
	}
}

// LogContext returns the build information in a log context format.
func (b BuildInfo) LogContext() map[string]interface{} {
	return map[string]interface{}{
		"version":     b.Version,
		"commit_hash": b.CommitHash,
		"build_date":  b.BuildDate,
		"go_version":  b.GoVersion,
		"os":          b.Os,
		"arch":        b.Arch,
		"compiler":    b.Compiler,
	}
}
