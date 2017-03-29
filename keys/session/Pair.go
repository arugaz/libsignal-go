package session

import (
	"github.com/RadicalApp/libsignal-protocol-go/ecc"
	"github.com/RadicalApp/libsignal-protocol-go/keys/chain"
	"github.com/RadicalApp/libsignal-protocol-go/keys/message"
)

// RootKeyable is an interface for all root key implementations that are part of
// a session keypair.
type RootKeyable interface {
	Bytes() []byte
	CreateChain(theirRatchetKey ecc.ECPublicKeyable, ourRatchetKey *ecc.ECKeyPair) (*KeyPair, error)
}

// ChainKeyable is an interface for all chain key implementations that are part of
// a session keypair.
type ChainKeyable interface {
	Key() []byte
	Index() uint32
	NextKey() *chain.Key
	MessageKeys() *message.Keys
	Current() *chain.Key
}

// NewKeyPair returns a new session key pair that holds a root and chain key.
func NewKeyPair(rootKey RootKeyable, chainKey ChainKeyable) *KeyPair {
	keyPair := KeyPair{
		RootKey:  rootKey,
		ChainKey: chainKey,
	}

	return &keyPair
}

// KeyPair is a session key pair that holds a single root and chain key pair. These
// keys are ratcheted after every message sent and every message round trip.
type KeyPair struct {
	RootKey  RootKeyable
	ChainKey ChainKeyable
}
