package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// removeCmd.PersistentFlags().StringVar(&path, "path", "", "Path to the makefile")
	// removeCmd.PersistentFlags().StringVar(&alias, "alias", "", "Project alias")

	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all currently registered makefiles",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		commandList()
	},
}

func commandList() {
	aliases := viper.AllKeys()

	if len(aliases) == 0 {
		fmt.Println("No makefile registered. Run mek add {alias} to add new ones")
	}

	for _, alias := range aliases {
		path := viper.GetString(alias)

		if path != "" {
			fmt.Println(alias + ":" + path)
		}
	}
}
