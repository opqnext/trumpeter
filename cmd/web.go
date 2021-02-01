package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"trumpeter/web"
)

var webPort int

func init() {
	webCmd.PersistentFlags().IntVarP(&webPort, "port", "w", 8080, "web port output")
	rootCmd.AddCommand(webCmd)
	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//imageCmd.AddCommand(cmdTimes)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "this is web view by websocket",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("执行web服务: web : %v", args)
		web.Run()
	},
}
