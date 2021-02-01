package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"trumpeter/db"
)

var host string
var port int

func init() {
	collectCmd.PersistentFlags().StringVarP(&host, "host", "l", "localhost", "redis host addr")
	collectCmd.PersistentFlags().IntVarP(&port, "port", "p", 6379, "redis port")
	rootCmd.AddCommand(collectCmd)
	//rootCmd.MarkFlagRequired("host")
	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//imageCmd.AddCommand(cmdTimes)
}

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect data.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("执行采集的方法:collect host : %v", args)

		db.InitRedis(host, port)
		db.Pop()
	},
}
