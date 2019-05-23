// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis instead.
package travis

import testing "github.com/libp2p/go-libp2p-testing/ci/travis"

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.EnvVar instead.
type EnvVar = testing.EnvVar

const (
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.EnvVar instead.
	VarCI = testing.VarCI
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarTravis instead.
	VarTravis = testing.VarTravis
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarBranch instead.
	VarBranch = testing.VarBranch
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarBuildDir instead.
	VarBuildDir = testing.VarBuildDir
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarBuildId instead.
	VarBuildId = testing.VarBuildId
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarBuildNumber instead.
	VarBuildNumber = testing.VarBuildNumber
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarCommit instead.
	VarCommit = testing.VarCommit
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarCommitRange instead.
	VarCommitRange = testing.VarCommitRange
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarJobId instead.
	VarJobId = testing.VarJobId
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarJobNumber instead.
	VarJobNumber = testing.VarJobNumber
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarPullRequest instead.
	VarPullRequest = testing.VarPullRequest
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarSecureEnvVars instead.
	VarSecureEnvVars = testing.VarSecureEnvVars
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarRepoSlug instead.
	VarRepoSlug = testing.VarRepoSlug
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarOsName instead.
	VarOsName = testing.VarOsName
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarTag instead.
	VarTag = testing.VarTag
	// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.VarGoVersion instead.
	VarGoVersion = testing.VarGoVersion
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.IsRunning instead.
func IsRunning() bool {
	return testing.IsRunning()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.Env instead.
func Env(v EnvVar) string {
	return testing.Env(v)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.JobId instead.
func JobId() string {
	return testing.JobId()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/ci/travis.JobNumber instead.
func JobNumber() string {
	return testing.JobNumber()
}
