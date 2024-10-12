package tool_core

import (
	"fmt"
	"testing"
)

func TestCheckLocalCert(*testing.T) {
	filePath := "../cert/xxx.pem"
	if data, ok := CheckLocalCert(&filePath); !ok {
		panic("Test failed")
	} else {
		fmt.Println(*data)
	}

}

func TestCheckRemoteCert(*testing.T) {
	url := "www.example.com:443"
	if certs, ok := CheckRemoteCert(&url); !ok {
		panic("Test failed")
	} else {
		fmt.Println((*certs)[0])
	}
}
