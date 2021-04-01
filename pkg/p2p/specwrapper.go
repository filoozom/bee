package p2p

import (
	"context"
	"errors"
	"time"
)

var errUnexpected = errors.New("unexpected request while in light mode")

// WithDisconnectStreams will mutate the given spec and replace the handler with a always erroring one
func WithDisconnectStreams(spec ProtocolSpec) {
	for i := range spec.StreamSpecs {
		spec.StreamSpecs[i].Handler = func(c context.Context, p Peer, s Stream) error {
			return NewDisconnectError(errUnexpected)
		}
	}
}

// WithBlocklistStreams will mutate the given spec and replace the handler with a always erroring one
func WithBlocklistStreams(dur time.Duration, spec ProtocolSpec) {
	for i := range spec.StreamSpecs {
		spec.StreamSpecs[i].Handler = func(c context.Context, p Peer, s Stream) error {
			return NewBlockPeerError(dur, errUnexpected)
		}
	}
}