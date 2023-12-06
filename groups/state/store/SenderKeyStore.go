package store

import (
	"github.com/arugaz/libsignal/groups/state/record"
	"github.com/arugaz/libsignal/protocol"
)

type SenderKey interface {
	StoreSenderKey(senderKeyName *protocol.SenderKeyName, keyRecord *record.SenderKey)
	LoadSenderKey(senderKeyName *protocol.SenderKeyName) *record.SenderKey
}
