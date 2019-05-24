// Deprecated: use github.com/libp2p/go-libp2p-testing and subpackages instead.
package testutil

import (
	"context"
	"math/rand"
	"testing"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/test"

	"github.com/libp2p/go-libp2p-testing/etc"
	"github.com/libp2p/go-libp2p-testing/net"

	ma "github.com/multiformats/go-multiaddr"
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.ZeroLocalTCPAddress instead.
// Warning: it's impossible to alias a var in go. Writes to this var would have no longer
// have any effect, so it has been commented out to induce breakage for added safety.
// var ZeroLocalTCPAddress = tnet.ZeroLocalTCPAddress

// Deprecated: use github.com/libp2p/go-libp2p-core/test.RandTestKeyPair instead.
// Supply RSA as a key type to get an equivalent result.
func RandTestKeyPair(bits int) (ci.PrivKey, ci.PubKey, error) {
	return test.RandTestKeyPair(ci.RSA, bits)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/test.SeededTestKeyPair instead.
// Supply RSA as a key type, with 512 bits, to get an equivalent result.
func SeededTestKeyPair(seed int64) (ci.PrivKey, ci.PubKey, error) {
	return test.SeededTestKeyPair(ci.RSA, 512, seed)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/test.RandPeerID instead.
func RandPeerID() (peer.ID, error) {
	return test.RandPeerID()
}

// Deprecated: use github.com/libp2p/go-libp2p-core/test.RandPeerIDFatal instead.
func RandPeerIDFatal(t testing.TB) peer.ID {
	return test.RandPeerIDFatal(t)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.RandLocalTCPAddress instead.
func RandLocalTCPAddress() ma.Multiaddr {
	return tnet.RandLocalTCPAddress()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.PeerNetParams instead.
type PeerNetParams = tnet.PeerNetParams

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.RandPeerNetParamsOrFatal instead.
func RandPeerNetParamsOrFatal(t *testing.T) PeerNetParams {
	return tnet.RandPeerNetParamsOrFatal(t)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.RandPeerNetParams instead.
func RandPeerNetParams() (*PeerNetParams, error) {
	return tnet.RandPeerNetParams()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.Identity instead.
type Identity = tnet.Identity

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.RandIdentity instead.
func RandIdentity() (Identity, error) {
	return tnet.RandIdentity()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.RandIdentityOrFatal instead.
func RandIdentityOrFatal(t *testing.T) Identity {
	return tnet.RandIdentityOrFatal(t)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.NewIdentity instead.
func NewIdentity(ID peer.ID, addr ma.Multiaddr, privk ci.PrivKey, pubk ci.PubKey) Identity {
	return tnet.NewIdentity(ID, addr, privk, pubk)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/net.LatencyConfig instead.
type LatencyConfig = tnet.LatencyConfig

// Deprecated: use github.com/libp2p/go-libp2p-testing/etc.SeededRand instead.
var SeededRand = tetc.SeededRand

// Deprecated: use github.com/libp2p/go-libp2p-testing/etc.NewSeededRand instead.
func NewSeededRand(seed int64) *rand.Rand {
	return tetc.NewSeededRand(seed)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/etc.LockedRandSource instead.
type LockedRandSource = tetc.LockedRandSource

// Deprecated: use github.com/libp2p/go-libp2p-testing/etc.WaitFor instead.
func WaitFor(ctx context.Context, check func() error) error {
	return tetc.WaitFor(ctx, check)
}
