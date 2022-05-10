package config

import "os"

// CertificateAuthority is path to certificate authority
var CertificateAuthority string

// LoadFromEnvironment loads configuration parameter from environment
func LoadFromEnvironment(v *string, key string) {
	if fromEnv := os.Getenv(key); fromEnv != "" {
		*v = fromEnv
	}
}
