package tool_core

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

type CertInfo struct {
	Subject   string `json:"subject" yaml:"subject"`
	Issuer    string `json:"issuer" yaml:"issuer"`
	NotBefore string `json:"not_before" yaml:"not_before"`
	NotAfter  string `json:"not_after" yaml:"not_after"`
}

func CheckLocalCert(filePath *string) (*CertInfo, bool) {
	fileByte, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
		return &CertInfo{}, false
	}
	block, _ := pem.Decode(fileByte)
	if block == nil {
		log.Fatal("解析PEM文件失败")
		return &CertInfo{}, false
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
		return &CertInfo{}, false
	}
	return &CertInfo{
		Subject:   cert.Subject.String(),
		Issuer:    cert.Issuer.String(),
		NotBefore: cert.NotBefore.String(),
		NotAfter:  cert.NotAfter.String(),
	}, true
}

func CheckRemoteCert(url *string) (*[]CertInfo, bool) {
	var result []CertInfo
	conn, err := tls.Dial("tcp", *url, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
		return &[]CertInfo{}, false
	}
	defer func(conn *tls.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		result = append(result, CertInfo{
			Subject:   cert.Subject.String(),
			Issuer:    cert.Issuer.String(),
			NotBefore: cert.NotBefore.String(),
			NotAfter:  cert.NotAfter.String(),
		})
	}

	return &result, true
}
