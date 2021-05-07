package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示 trumpeter 当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trumpeter Static Site Generator v0.9 -- HEAD")
	},
}
