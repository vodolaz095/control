package commands

import (
	"crypto/tls"
	"crypto/x509"
	"gitflic.ru/vodolaz095/control/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var startAgentCommand = &cobra.Command{
	Use:     "agent",
	Short:   "Запустить программу в режиме агента",
	Long:    "Запустить программу в режиме агента",
	Example: " control agent",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Starting agent...")
		config.LoadFromEnvironment(&config.Server, "SERVER")

		config.LoadFromEnvironment(&config.ClientKey, "CLIENT_KEY")
		logger.Printf("Loading client key from %s...", config.ClientKey)

		config.LoadFromEnvironment(&config.ClientCert, "CLIENT_CERT")
		logger.Printf("Loading client certificate from %s...", config.ClientCert)

		cert, err := tls.LoadX509KeyPair(config.ClientCert, config.ClientKey)
		if err != nil {
			logger.Fatalf("%s : while loading key", err)
		}
		logger.Printf("Server certificate loaded!")

		log.Printf("Preparing to start client connecting to %s with certificate from %s",
			config.Server, config.ClientCert,
		)

		config.LoadFromEnvironment(&config.CertificateAuthority, "CA")
		logger.Printf("Loading certificate authority from %s...", config.CertificateAuthority)
		certificateAuthorityData, err := ioutil.ReadFile(config.CertificateAuthority)
		if err != nil {
			logger.Fatalf("%s : while loading certificate authority from %s", err, config.CertificateAuthority)
		}
		authority := x509.NewCertPool()
		authority.AppendCertsFromPEM(certificateAuthorityData)

		log.Printf("Preparing to start client connecting to %s with certificate from %s",
			config.Server, config.ClientCert,
		)
		cfg := &tls.Config{
			ServerName:   "queue2.vodolaz095.life",
			RootCAs:      authority,
			Certificates: []tls.Certificate{cert},
		}
		conn, err := tls.Dial("tcp", config.Server, cfg)
		if err != nil {
			log.Fatalf("%s : while dialing %s", err, config.Server)
		}
		data, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatalf("%s : while reading all data from server", err)
		}
		logger.Printf("Response: %s", string(data))
	},
}
