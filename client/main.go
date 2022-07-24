package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"net/http"
)

func MakeDownloadRequest(useTls bool, url string) (*http.Response, error) {

	var client *http.Client
	if useTls {
		fmt.Println("use tls")
		/*caCert, err := ioutil.ReadFile("/data/the-mesh-for-data/ca-cert.pem")
		if err != nil {
			return nil, err
		}
		caCertPool, _ := x509.SystemCertPool()
		caCertPool.AppendCertsFromPEM(caCert)*/
		var caCertPool *x509.CertPool
		certs, _ := tls.LoadX509KeyPair("/tmp/certs/client-cert.pem", "/tmp/certs/client-key.pem")

		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      caCertPool,
					Certificates: []tls.Certificate{certs},
				},
			},
		}
	} else {
		client = http.DefaultClient
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	useTls := flag.Bool("useTls", false, "use tls")
	url := flag.String("url", "https://www.google.com/", "url")
	flag.Parse()
	resp, err := MakeDownloadRequest(*useTls, *url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.Status)
}
