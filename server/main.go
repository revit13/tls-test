package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})
	loadedCertServer, _ := tls.LoadX509KeyPair("/tmp/certs/server-cert.pem", "/tmp/certs/server-key.pem")
	var CACertPool *x509.CertPool
	/*CACertPool := x509.NewCertPool()
	cert, err := ioutil.ReadFile("/tmp/certs/ca-cert.pem")
	if err != nil {
		log.Fatalf("Couldn't load file", err)
	}
	CACertPool.AppendCertsFromPEM(cert)*/
	cfg := &tls.Config{
		Certificates: []tls.Certificate{loadedCertServer},
		// configure mutual TLS
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  CACertPool,
		MinVersion: tls.VersionTLS13,
	}
	srv := &http.Server{
		Addr:         ":8443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func main() {
	server()
}
