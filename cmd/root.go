package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ip-range-processor",
		Short: "ip-range-processor",
		Long:  `ip-range-processor`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
