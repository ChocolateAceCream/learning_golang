package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

const url = "https://121.41.170.188:8000"

var httpVersion = flag.Int("version", 2, "HTTP version")

func main() {

	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("server-cert.pem")

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(pemServerCA)

	// Create the credentials and return it
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	flag.Parse()
	client := &http.Client{}

	// Use the proper transport in the client
	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	// Perform the request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf(
		"Got response %d: %s %s\n",
		resp.StatusCode, resp.Proto, string(body))
}
