package commands

import (
	"fmt"
	"ssl_tool/tool_core"

	"github.com/spf13/cobra"
)

func NewSslToolCmd() *cobra.Command {
	var url, filePath string
	rootCmd := &cobra.Command{
		Use:   "ssl-tool",
		Short: "SSL 工具检测证书的信息",
		Run: func(cmd *cobra.Command, args []string) {
			if url != "" {
				// 检测远端证书
				tool_core.CheckRemoteCert(&url)
			} else if filePath != "" {
				// 检测本地证书
				tool_core.CheckLocalCert(&filePath)
			} else {
				fmt.Println(cmd.Help())
			}

		},
	}
	rootCmd.Flags().StringVarP(&url, "internet", "i", "", "指定要检测的域名")
	rootCmd.Flags().StringVarP(&filePath, "local", "l", "", "指定要检测的本地证书文件夹子")
	return rootCmd
}
