// Deprecated: use github.com/libp2p/go-libp2p-testing/ci instead.
package ci

import testing "github.com/libp2p/go-libp2p-testing/ci"

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.EnvVar instead.
type EnvVar = testing.EnvVar

const (
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.VarCI instead.
	VarCI = testing.VarCI
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.VarNoFuse instead.
	VarNoFuse = testing.VarNoFuse
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.VarVerbose instead.
	VarVerbose = testing.VarVerbose
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.IsRunning instead.
func IsRunning() bool {
	return testing.IsRunning()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.EnvVar instead.
func Env(v EnvVar) string {
	return testing.Env(v)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.NoFuse instead.
func NoFuse() bool {
	return testing.NoFuse()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci.Verbose instead.
func Verbose() bool {
	return testing.Verbose()
}
