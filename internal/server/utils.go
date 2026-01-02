package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func SetupTLS() {
	pool := x509.NewCertPool()
	certFile := "/ca-certificates.crt"
	fi, err := os.ReadFile(certFile)
	if err != nil {
		slog.Warn(fmt.Sprintf("Could not open %s for reading CAs", certFile))
	} else {
		ok := pool.AppendCertsFromPEM(fi)
		if !ok {
			slog.Warn("Certificates were not parsed correctly")
		}
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{RootCAs: pool},
			},
		}
		*http.DefaultClient = *client
	}
}
