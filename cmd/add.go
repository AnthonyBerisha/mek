package cmd

import (
	"fmt"
	"os"

	// "os"
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
		main()
	},
}

func main() {
	// fmt.Println(path, alias)

	if _, err := os.Stat("~/.config/mek"); os.IsNotExist(err) {
		err := os.Mkdir("~/.config/mek", os.ModeDir)

		if err != nil {
			fmt.Println("Error creating folder mek", err)
		}

	}

	// f, err := os.Create("~/.config/mek/mek.yaml")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err, f)
	// }
}
