package tool_core

import "testing"

func TestCheckLocalCert(*testing.T) {
	filePath := "../cert/xxx.pem"
	if _, ok := CheckLocalCert(&filePath); !ok {
		panic("Test failed")
	}
}

func TestCheckRemoteCert(*testing.T) {
	url := "https://www.baidu.com"
	if _, ok := CheckRemoteCert(&url); !ok {
		panic("Test failed")
	}
}
