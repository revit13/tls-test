# reference: https://www.youtube.com/watch?v=7YgaZIFn7mY
# generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4656 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=dd/ST=dd/O=Internet Widgits Pty Ltd" -nodes
openssl x509 -in ca-cert.pem -noout -text

# generate the server private key and certificate signing request
openssl req  -newkey rsa:4656  -keyout server-key.pem -out server-req.pem -subj "/C=dd/ST=dd/O=webse" -nodes
# use CA's private key to sign server's CSR and get back the sertificate
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
# Server's signed cerificate
openssl x509 -in server-cert.pem -noout -text


# generate the client private key and certificate signing request
openssl req  -newkey rsa:4656  -keyout client-key.pem -out client-req.pem -subj "/C=dd/ST=dd/O=webse" -nodes

# use CA's private key to sign server's CSR and get back the sertificate
openssl x509 -req -in client-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

mkdir -p /tmp/certs/
cp *.pem /tmp/certs/

sudo cp ca-cert.pem /etc/ssl/certs
