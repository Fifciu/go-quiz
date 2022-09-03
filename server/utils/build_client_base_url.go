package utils

import (
	"fmt"
	"os"
)

func BuildClientBaseUrl() string {
	protocol := os.Getenv("client_protocol")
	host := os.Getenv("client_host")
	port := os.Getenv("client_port")

	return fmt.Sprintf("%s://%s:%s", protocol, host, port)
}
