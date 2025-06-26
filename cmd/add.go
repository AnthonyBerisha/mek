package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	if path == "" {
		println("Path is required")
		return
	}

	if alias == "" {
		println("Alias is required")
		return
	}

	addMakefileToConfig(path, alias)

	// println(pathErr.Error)

}

func addMakefileToConfig(path string, alias string) error {
	var err error

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("path is incorrect")
	}

	viper.Set(alias, path)

	viper.WriteConfig()

	println(viper.ReadInConfig())

	// Open config file
	// os.OpenFile(configFilePath)
	// Check path is not in config file
	// Check alias is not in config file
	// Write alias=path
	// Load into memory and split at the = symbole

	return err
}
