package server

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/pkg/errors"
	"os"
)

// NewTLSConfig returns a tls.Config for the given Config
func NewTLSConfig(cfg TLSConfig) (*tls.Config, error) {
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if !strSliceContains(config.NextProtos, "http/1.1") {
		// Enable HTTP/1.1
		config.NextProtos = append(config.NextProtos, "http/1.1")
	}

	if !strSliceContains(config.NextProtos, "h2") {
		// Enable HTTP/2
		config.NextProtos = append([]string{"h2"}, config.NextProtos...)
	}

	if !cfg.GetEnabled() {
		return config, nil
	}

	certPEMBlock, err := os.ReadFile(cfg.CertFile)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"tls: failed to load certificate",
		)
	}

	keyPEMBlock, err := os.ReadFile(cfg.KeyFile)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"tls: failed to load private key",
		)
	}

	config.ClientCAs = x509.NewCertPool()

	if len(cfg.RootCAFile) > 0 {
		rootcaPEMBlock, err := os.ReadFile(cfg.RootCAFile)
		if err != nil {
			return nil, errors.Wrap(
				err,
				"tls: failed to load RootCA certificates",
			)
		}

		if !config.ClientCAs.AppendCertsFromPEM(rootcaPEMBlock) {
			return nil, errors.Wrap(
				err,
				"tls: failed to append RootCA certificates",
			)
		}
	}

	c, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"tls: failed creating key pair",
		)
	}

	config.Certificates = append(config.Certificates, c)

	return config, nil
}

func strSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
