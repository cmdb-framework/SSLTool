package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"ssl_tool/tool_core"
)

func NewSslToolCmd() *cobra.Command {
	var url, filePath, output string
	rootCmd := &cobra.Command{
		Use:   "ssl-tool",
		Short: "SSL 工具检测证书的信息",
		Run: func(cmd *cobra.Command, args []string) {
			if url != "" {
				if certs, ok := tool_core.CheckRemoteCert(&url); !ok {
					fmt.Println("检测远程证书失败")
				} else {
					consolePrint(&(*certs)[0], &output)
				}
			} else if filePath != "" {
				if cert, ok := tool_core.CheckLocalCert(&filePath); !ok {
					fmt.Println("检测本地证书失败")
				} else {
					consolePrint(cert, &output)
				}
			} else {
				fmt.Println(cmd.Help())
			}
		},
	}
	rootCmd.Flags().StringVarP(&url, "internet", "i", "", "指定要检测的域名")
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "指定要检测的本地证书文件夹子")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "输出的格式，目前支持json")
	return rootCmd
}

func consolePrint(c *tool_core.CertInfo, o *string) {
	switch *o {
	case "json":
		if marshal, err := json.Marshal(*c); err != nil {
			return
		} else {
			fmt.Println(string(marshal))
		}
	case "yaml":
		if marshal, err := yaml.Marshal(*c); err != nil {
			return
		} else {
			fmt.Println(string(marshal))
		}
	default:
		fmt.Println("开始时间:", c.NotBefore)
		fmt.Println("结束结束", c.NotAfter)
		fmt.Println("颁发者:", c.Issuer)
		fmt.Println("主题:", c.Subject)
	}
}
