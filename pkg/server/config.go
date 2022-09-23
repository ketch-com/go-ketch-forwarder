package server

import (
	"fmt"
	"os"
	"strconv"
)

func getEnvOrDefault(key string, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		v = defaultValue
	}
	return v
}

func getEnvOrDefaultAsUInt(key string, defaultValue uint) uint {
	v, err := strconv.ParseUint(getEnvOrDefault(key, fmt.Sprint(defaultValue)), 10, 64)
	if err != nil {
		return defaultValue
	}
	return uint(v)
}

func NewConfig() Config {
	return Config{
		Bind:   getEnvOrDefault("KETCH_BIND", "0.0.0.0"),
		Listen: getEnvOrDefaultAsUInt("KETCH_LISTEN", 5000),
		TLS: TLSConfig{
			CertFile:   getEnvOrDefault("KETCH_TLS_CERT_FILE", ""),
			KeyFile:    getEnvOrDefault("KETCH_TLS_KEY_FILE", ""),
			RootCAFile: getEnvOrDefault("KETCH_TLS_ROOTCA_FILE", ""),
		},
		Username: getEnvOrDefault("KETCH_USER_NAME", ""),
		Password: getEnvOrDefault("KETCH_USER_PASSWORD", ""),
	}
}

type Config struct {
	Bind     string
	Listen   uint
	TLS      TLSConfig
	Username string
	Password string
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Bind, c.Listen)
}

type TLSConfig struct {
	CertFile   string
	KeyFile    string
	RootCAFile string
}

func (c TLSConfig) GetEnabled() bool {
	return len(c.CertFile) > 0 || len(c.KeyFile) > 0
}
