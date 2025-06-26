package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete the configuration folder and file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		commandClean()
	},
}

func commandClean() {
	var folderPath = getFolderPath()

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		println("No folder found", err.Error())
	}

	// Attempt to remove the folder
	err := os.RemoveAll(folderPath)

	if err != nil {
		log.Fatal("Error removing folder at:", folderPath, err.Error())
	} else {
		println("Folder removed successfully at path:", folderPath)
	}
}
