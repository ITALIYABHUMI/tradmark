package cmd

import (
	"github.com/spf13/cobra"
)

var unzipCmd = &cobra.Command{
	Use:     "unzip",
	Aliases: []string{"unzip"},
	Short:   "unzip file",
	Long:    "unzip file or link in the other folder",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Unzip(args[0])
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)
}
