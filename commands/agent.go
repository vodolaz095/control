package commands

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"gitflic.ru/vodolaz095/control/agent"
	"io/ioutil"
	"net"
	"time"

	"gitflic.ru/vodolaz095/control/config"
	"gitflic.ru/vodolaz095/control/pb"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var startAgentCommand = &cobra.Command{
	Use:     "agent",
	Short:   "Запустить программу в режиме агента",
	Long:    "Запустить программу в режиме агента",
	Example: " control agent",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Starting agent...")
		config.LoadFromEnvironment(&config.Server, "SERVER_ADDR")
		config.LoadFromEnvironment(&config.ServerName, "SERVER_NAME")

		config.LoadFromEnvironment(&config.ClientKey, "CLIENT_KEY")
		logger.Printf("Loading client key from %s...", config.ClientKey)

		config.LoadFromEnvironment(&config.ClientCert, "CLIENT_CERT")
		logger.Printf("Loading client certificate from %s...", config.ClientCert)

		cert, err := tls.LoadX509KeyPair(config.ClientCert, config.ClientKey)
		if err != nil {
			logger.Fatalf("%s : while loading key", err)
		}
		logger.Printf("Server certificate loaded!")

		logger.Printf("Preparing to start client connecting to %s with certificate from %s",
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

		logger.Printf("Preparing to start client connecting to %s with certificate from %s",
			config.Server, config.ClientCert,
		)
		cfg := &tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ServerName:   config.ServerName,
			RootCAs:      authority,
			Certificates: []tls.Certificate{cert},
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, config.Server,
			grpc.WithContextDialer(func(ctx context.Context, s string) (tlsCon net.Conn, err error) {
				logger.Printf("Dialing with TLS: %s", s)
				when, ok := ctx.Deadline()
				if ok {
					dialer := net.Dialer{Deadline: when}
					tlsCon, err = tls.DialWithDialer(&dialer, "tcp", s, cfg)
					return
				}
				tlsCon, err = tls.Dial("tcp", s, cfg)
				return
			}),
			//grpc.WithAuthority("queue2.vodolaz095.life"), // means nothing?
			grpc.WithTransportCredentials(credentials.NewTLS(cfg)),
			grpc.WithBlock(),
		)
		if err != nil {
			logger.Fatalf("%s : while dialing %s", err, config.Server)
		}
		logger.Printf("Connection established to %s", conn.Target())

		client := pb.NewControlClient(conn)
		resp, err := client.GetLine(ctx, &pb.StringRequest{Data: "something"})
		if err != nil {
			logger.Fatalf("%s : while getting response for getLine", err)
		}
		logger.Printf("Response: %s", resp.String())

		task, err := client.GetTaskByName(ctx, &pb.StringRequest{Data: "something"})
		if err != nil {
			logger.Fatalf("%s : while getting task something", err)
		}
		logger.Printf("Response: %s", task.String())

		go agent.StartReportingTelemetry(client, logger)
		agent.StartTaskLoop("agent", client, logger)

		err = conn.Close()
		if err != nil {
			logger.Fatalf("%s : while closing connection", err)
		}
	},
}

// Good read
// https://stackoverflow.com/questions/71976729/golang-tls-handshake-error-first-record-does-not-look-like-a-tls-handshake
