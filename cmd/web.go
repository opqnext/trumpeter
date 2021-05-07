package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"trumpeter/web"
)

var webPort int

func init() {
	//webCmd.PersistentFlags().IntVarP(&webPort, "port", "w", 1215, "web port output")
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "运行一个基于websocket的可视化界面", //(this is web view by websocket)
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("执行web: %s\n", "http://localhost:1215")
		webPort = 1215
		web.Run(webPort)
	},
}
