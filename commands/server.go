package commands

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gitflic.ru/vodolaz095/control/config"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var startServerCommand = &cobra.Command{
	Use:     "server",
	Short:   "Запустить программу в режиме сервера",
	Long:    "Запустить программу в режиме сервера",
	Example: " control server",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		config.LoadFromEnvironment(&config.ADDR, "ADDR")
		config.LoadFromEnvironment(&config.PORT, "PORT")
		log.Printf("Preparing to start server on %s:%s...", config.ADDR, config.PORT)

		config.LoadFromEnvironment(&config.ServerKey, "SERVER_KEY")
		logger.Printf("Loading server key from %s...", config.ServerKey)
		config.LoadFromEnvironment(&config.ServerCert, "SERVER_CERT")
		logger.Printf("Loading server certificate from %s...", config.ServerCert)
		cert, err := tls.LoadX509KeyPair(config.ServerCert, config.ServerKey)
		if err != nil {
			logger.Fatalf("%s : while loading key", err)
		}
		logger.Printf("Server certificate loaded!")

		config.LoadFromEnvironment(&config.CertificateAuthority, "CA")
		logger.Printf("Loading certificate authority from %s...", config.CertificateAuthority)
		certificateAuthorityData, err := ioutil.ReadFile(config.CertificateAuthority)
		if err != nil {
			logger.Fatalf("%s : while loading certificate authority from %s", err, config.CertificateAuthority)
		}
		authority := x509.NewCertPool()
		authority.AppendCertsFromPEM(certificateAuthorityData)

		logger.Printf("Preparing to start server on %s:%s...", config.ADDR, config.PORT)
		cfg := &tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    authority,
			Certificates: []tls.Certificate{cert},
		}
		listener, err := tls.Listen("tcp",
			fmt.Sprintf("%s:%s", config.ADDR, config.PORT),
			cfg,
		)
		if err != nil {
			logger.Fatalf("%s : while starting server on %s:%s", err, config.ADDR, config.PORT)
		}
		logger.Printf("Listening for incoming connections...")
		for {
			con, err := listener.Accept()
			if err != nil {
				logger.Printf("%s : while accepting connection", err)
				continue
			}
			logger.Printf("Connection from %s...", con.RemoteAddr().String())
			_, err = fmt.Fprintln(con, "HELLO")
			if err != nil {
				logger.Printf("%s : while accepting connection", err)
				con.Close()
				continue
			}
			err = con.Close()
			if err != nil {
				logger.Printf("%s : while accepting connection", err)
			}
		}
	},
}
