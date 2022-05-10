package commands

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"

	"gitflic.ru/vodolaz095/control/config"
	"gitflic.ru/vodolaz095/control/pb"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

type RPC struct {
	pb.UnimplementedControlServer
	Logger *log.Logger
}

func (rpc *RPC) GetLine(ctx context.Context, in *pb.StringRequest) (*pb.StringResponse, error) {
	rv := in.Data
	rpc.Logger.Printf("Receive: %v", rv)
	return &pb.StringResponse{Data: rv}, nil
}

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

		rpc := new(RPC)
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
