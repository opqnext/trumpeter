package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "trumpeter",
	Short: "trumpeter is a young and talented bag catcher",
	Long:  `trumpeter is a young and talented bag catcher`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		fmt.Printf("11=== %v \n", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
