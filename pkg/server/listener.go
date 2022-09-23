package server

import (
	"crypto/tls"
	"github.com/ketch-com/go-ketch-forwarder/version"
	"github.com/pkg/errors"
	"log"
	"net"
)

// NewListener returns a net.Listener for the given config
func NewListener(config Config) (net.Listener, error) {
	if !config.TLS.GetEnabled() {
		return nil, errors.New("TLS is required")
	}

	// Start listening
	listener, err := net.Listen("tcp", config.Addr())
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to listen on %s",
			config.Addr(),
		)
	}

	// If TLS is not enabled, return the listener as it is
	if !config.TLS.GetEnabled() {
		return listener, nil
	}

	// Since TLS is enabled, get the tls.Config and return a crypto/tls listener
	cfg, err := NewTLSConfig(config.TLS)
	if err != nil {
		_ = listener.Close()

		return nil, errors.Wrap(
			err,
			"failed to get server TLS config",
		)
	}

	log.Printf("⚡️ Ketch Event Forwarder %s listening on port %s\n", version.Version, config.Addr())

	return tls.NewListener(listener, cfg), nil
}
