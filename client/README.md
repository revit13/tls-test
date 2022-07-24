# cleint side tls

Usage: go run main.go -useTls=<bool> -url=<url>

useTls: Specifies whether tls is used. Default value is false.
url: the url for testing. Default value is "http://google.com".

example:

go run main.go -useTls=true -url="https://localhost:8443"
