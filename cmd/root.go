package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "UnzipAndStoreData",
	Short: "UnzipAndStoreData is a cli tool for performing file system operations",
	Long:  "UnzipAndStoreData is a cli tool for performing basic file system operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing UnzipAndStoreData '%s'\n", err)
		os.Exit(1)
	}
}
