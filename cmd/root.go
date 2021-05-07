package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "trumpeter",
	Short: "trumpeter 是一款用于实时采集、展示网络请求的抓包小工具",
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//
	//	fmt.Printf("11=== %v \n", args)
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
