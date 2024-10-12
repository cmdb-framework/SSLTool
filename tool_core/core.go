package tool_core

import (
	"crypto/x509"
	"encoding/pem"
	// "fmt"
	"log"
	"os"
	// "strings"
)

type CertInfo struct {
	Subject   string
	Issuer    string
	NotBefore string
	NotAfter  string
	Raw       string
}

func pemCheck(filePath *string) (*CertInfo, bool) {
	pemData, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
		return &CertInfo{}, false
	}

	block, _ := pem.Decode(pemData)
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

func CheckLocalCert(filePath *string) (*CertInfo, bool) {
	// fileTypeArray := strings.Split(*filePath, ".")
	// switch fileType := fileTypeArray[len(fileTypeArray)-1]; fileType {
	// case "pem":
	// 	fmt.Println("PEM文件")
	// case "crt":
	// 	fmt.Println("CRT文件")
	// case "cer":
	// 	fmt.Println("CER文件")
	// default:
	// 	log.Fatal("文件类型不支持")
	// }
	pemData, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
		return &CertInfo{}, false
	}

	block, _ := pem.Decode(pemData)
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

func CheckRemoteCert(usrl *string) (*CertInfo, bool) {
	println("CheckRemoteCert")
	return &CertInfo{}, true
}
