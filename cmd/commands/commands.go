package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"ssl_tool/tool_core"
)

func NewSslToolCmd() *cobra.Command {
	var url, filePath string
	rootCmd := &cobra.Command{
		Use:   "ssl-tool",
		Short: "SSL 工具检测证书的信息",
		Run: func(cmd *cobra.Command, args []string) {
			if url != "" {
				if certs, ok := tool_core.CheckRemoteCert(&url); !ok {
					fmt.Println("检测远程证书失败")
				} else {
					cnPrint(&(*certs)[0])
				}
			} else if filePath != "" {
				if cert, ok := tool_core.CheckLocalCert(&filePath); !ok {
					fmt.Println("检测本地证书失败")
				} else {
					cnPrint(cert)
				}
			} else {
				fmt.Println(cmd.Help())
			}
		},
	}
	rootCmd.Flags().StringVarP(&url, "internet", "i", "", "指定要检测的域名")
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "指定要检测的本地证书文件夹子")
	return rootCmd
}

func cnPrint(c *tool_core.CertInfo) {
	fmt.Println("开始时间:", c.NotBefore)
	fmt.Println("结束结束", c.NotAfter)
	fmt.Println("颁发者:", c.Issuer)
	fmt.Println("主题:", c.Subject)
}
