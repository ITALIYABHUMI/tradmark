package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var unzipCmd = &cobra.Command{
	Use:     "unzip",
	Aliases: []string{"unzip"},
	Short:   "unzip file",
	Long:    "unzip file or link in the other folder",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ExtractAndStoreFileData(args[0])
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)
}

func ExtractAndStoreFileData(file string) string {
	if strings.HasSuffix(file, ".zip") {
		filePath, err := Unzip(file)
		if err != nil {
			return fmt.Sprintf("failed to covert data to Json: %v", err)
		}

		if err = ProcessFile(filePath); err != nil {
			return fmt.Sprintf("failed to covert data to Json: %v", err)
		}
		log.Println("XML data inserted successfully")

	} else if strings.HasSuffix(file, ".xml") {

		if err := ProcessFile(file); err != nil {
			return fmt.Sprintf("failed to covert data to Json: %v", err)
		}
		log.Println("XML data inserted successfully")
	}

	return "File processed successfully"
}
