package cmd

import (
	"github.com/spf13/cobra"
)

var (
	path  string
	alias string
)

func init() {
	addCmd.PersistentFlags().StringVar(&path, "path", "", "Path to the makefile")
	addCmd.PersistentFlags().StringVar(&alias, "alias", "", "Project alias")

	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a project's makefile to the list of known makefiles",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		commandAdd()
	},
}

func commandAdd() {
	err := checkInit()

	if condition := err != nil; condition {
		println("Error checking init:", err.Error())
		return
	}

	if path == "" {
		println("Path is required")
		return
	}

	if alias == "" {
		println("Alias is required")
		return
	}

}
