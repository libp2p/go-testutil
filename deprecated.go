package testutil

import (
	"context"
	"math/rand"
	"testing"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	tcrypto "github.com/libp2p/go-libp2p-core/crypto/test"
	peer "github.com/libp2p/go-libp2p-core/peer"
	tpeer "github.com/libp2p/go-libp2p-core/peer/test"
	"github.com/libp2p/go-libp2p-testing/helper"

	ma "github.com/multiformats/go-multiaddr"
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.ZeroLocalTCPAddress instead.
var ZeroLocalTCPAddress = helper.ZeroLocalTCPAddress

// Deprecated: use github.com/libp2p/go-libp2p-core/crypto/test.RandTestKeyPair instead.
// Supply RSA as a key type to get an equivalent result.
func RandTestKeyPair(bits int) (ci.PrivKey, ci.PubKey, error) {
	return tcrypto.RandTestKeyPair(ci.RSA, bits)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/crypto/test.SeededTestKeyPair instead.
// Supply RSA as a key type, with 512 bits, to get an equivalent result.
func SeededTestKeyPair(seed int64) (ci.PrivKey, ci.PubKey, error) {
	return tcrypto.SeededTestKeyPair(ci.RSA, 512, seed)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/peer/test.RandPeerID instead.
func RandPeerID() (peer.ID, error) {
	return tpeer.RandPeerID()
}

// Deprecated: use github.com/libp2p/go-libp2p-core/peer/test.RandPeerIDFatal instead.
func RandPeerIDFatal(t testing.TB) peer.ID {
	return tpeer.RandPeerIDFatal(t)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.RandLocalTCPAddress instead.
func RandLocalTCPAddress() ma.Multiaddr {
	return helper.RandLocalTCPAddress()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.PeerNetParams instead.
type PeerNetParams = helper.PeerNetParams

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.RandPeerNetParamsOrFatal instead.
func RandPeerNetParamsOrFatal(t *testing.T) PeerNetParams {
	return helper.RandPeerNetParamsOrFatal(t)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.RandPeerNetParams instead.
func RandPeerNetParams() (*PeerNetParams, error) {
	return helper.RandPeerNetParams()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.Identity instead.
type Identity = helper.Identity

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.RandIdentity instead.
func RandIdentity() (Identity, error) {
	return helper.RandIdentity()
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.RandIdentityOrFatal instead.
func RandIdentityOrFatal(t *testing.T) Identity {
	return helper.RandIdentityOrFatal(t)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.NewIdentity instead.
func NewIdentity(ID peer.ID, addr ma.Multiaddr, privk ci.PrivKey, pubk ci.PubKey) Identity {
	return helper.NewIdentity(ID, addr, privk, pubk)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.LatencyConfig instead.
type LatencyConfig = helper.LatencyConfig

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.SeededRand instead.
var SeededRand = helper.SeededRand

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.NewSeededRand instead.
func NewSeededRand(seed int64) *rand.Rand {
	return helper.NewSeededRand(seed)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.LockedRandSource instead.
type LockedRandSource = helper.LockedRandSource

// Deprecated: use github.com/libp2p/go-libp2p-testing/helper.WaitFor instead.
func WaitFor(ctx context.Context, check func() error) error {
	return helper.WaitFor(ctx, check)
}
