package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "mek",
		Short: "mek is a CLI tool to orchestrate your makefiles across different projects",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			alias := args[0]

			target := args[1]

			if target == "" {
				fmt.Printf("Missing makefile target")
				return
			}

			resolved := viper.GetString(alias)
			if resolved == "" {
				fmt.Printf("Unknown alias: %s\n", alias)
				return
			}

			dir := filepath.Dir(resolved)

			execCmd := exec.Command("make", target)
			execCmd.Dir = dir
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr
			execCmd.Stdin = os.Stdin

			if err := execCmd.Run(); err != nil {
				fmt.Println("Error running command:", err)
			}
		},
	}
)

// Executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/mek/mek.yaml)")
	// rootCmd.PersistentFlags().StringP("author", "a", "Anthony Berisha", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "Anthony Berisha berishaanthony@gmail.com")
	// viper.SetDefault("license", "MIT")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

func initConfig() {

	home, _ := os.UserHomeDir()

	viper.AddConfigPath(filepath.Join(home, ".config/mek"))

	viper.SetConfigName("mek")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		// Do nothing now ?
	}

	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := os.UserHomeDir()
	// 	cobra.CheckErr(err)

	// 	// Search config in home directory with name ".cobra" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName(".cobra")
	// }

	// viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}
