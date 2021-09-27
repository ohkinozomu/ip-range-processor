package cmd

import (
	"log"

	"github.com/ohkinozomu/ip-range-processor/pkg/process"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringP("input", "i", "", "input")
	runCmd.PersistentFlags().StringP("output", "o", "", "output")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run ip-range-processor",
	Long:  `run ip-range-processor`,
	Run: func(cmd *cobra.Command, args []string) {
		i, err := cmd.PersistentFlags().GetString("input")
		if err != nil {
			panic(err)
		}

		o, err := cmd.PersistentFlags().GetString("output")
		if err != nil {
			panic(err)
		}

		if i == "" || o == "" {
			log.Fatal("Set --input and --output")
		}

		if i == "datadog-synthetics" && o == "terraform-aws-waf" {
			err := process.Process()
			if err != nil {
				panic(err)
			}
		} else {
			log.Printf("Unsupported: %v %v", i, o)
		}
	},
}
