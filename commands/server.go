package commands

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"gitflic.ru/vodolaz095/control/config"
	"gitflic.ru/vodolaz095/control/pb"
	"gitflic.ru/vodolaz095/control/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/spf13/cobra"
)

// good read
// https://stackoverflow.com/a/43606908/1885921

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
			GetConfigForClient: func(info *tls.ClientHelloInfo) (*tls.Config, error) {
				logger.Printf("Client uses SNI to talk to %s...", info.ServerName)
				return nil, nil
			},
			VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) (err error) {
				var clientCert *x509.Certificate
				for _, oneOfClientCerts := range rawCerts {
					clientCert, err = x509.ParseCertificate(oneOfClientCerts)
					if err != nil {
						return
					}
					logger.Printf("Client certificate subject %s", clientCert.Subject.String())
					if clientCert.Subject.CommonName != "snow.local" {
						return fmt.Errorf("only snow.local is trusted")
					}
				}
				for i, chains := range verifiedChains {
					for j, chain := range chains {
						logger.Printf("Peer chain %v:%v subject %s", i, j, chain.Subject.String())
					}
				}
				return nil
			},
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

		rpc := new(server.RPC)
		rpc.Logger = logger

		s := grpc.NewServer(grpc.Creds(credentials.NewTLS(cfg)))
		pb.RegisterControlServer(s, rpc)

		err = s.Serve(listener)
		if err != nil {
			log.Fatalf("%s : while starting grpc server on %s:%s", err, config.ADDR, config.PORT)
		}

		//for {
		//	con, err := listener.Accept()
		//	if err != nil {
		//		logger.Printf("%s : while accepting connection", err)
		//		continue
		//	}
		//	logger.Printf("Connection from %s...", con.RemoteAddr().String())
		//	// https://stackoverflow.com/a/43606908/1885921
		//	err = con.(*tls.Conn).Handshake()
		//	if err != nil {
		//		logger.Printf("%s : while handshaking connection", err)
		//		continue
		//	}
		//
		//	logger.Printf("Handshake complete - %t", con.(*tls.Conn).ConnectionState().HandshakeComplete)
		//
		//	for _, cert := range con.(*tls.Conn).ConnectionState().PeerCertificates {
		//		logger.Printf("PeerCertificate valid for %s", cert.Subject.String())
		//	}
		//	logger.Printf("Verified chains...")
		//	for _, chain := range con.(*tls.Conn).ConnectionState().VerifiedChains {
		//		for _, cert := range chain {
		//			logger.Printf("Client cert valid for %s", cert.Subject.String())
		//		}
		//	}
		//
		//	_, err = fmt.Fprintln(con, "HELLO")
		//	if err != nil {
		//		logger.Printf("%s : while accepting connection", err)
		//		con.Close()
		//		continue
		//	}
		//	err = con.Close()
		//	if err != nil {
		//		logger.Printf("%s : while accepting connection", err)
		//	}
		//}
	},
}
