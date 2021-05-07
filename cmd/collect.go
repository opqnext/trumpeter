package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"trumpeter/db"
)

var host string
var port int
var key string

func init() {
	collectCmd.PersistentFlags().IntVarP(&port, "port", "p", 6379, "redis 端口号")
	collectCmd.PersistentFlags().StringVarP(&host, "host", "l", "localhost", "redis 地址")
	collectCmd.PersistentFlags().StringVarP(&key, "key", "k", "trumpeter:request:line", "redis 用于采集数据的key值")
	rootCmd.AddCommand(collectCmd)
}

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "使用Redis列表（lists）方式采集一次网络请求的数据",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("执行采集的方法:collect host : %v", args)

		db.InitRedis(host, port)
		db.Pop(key)
	},
}
