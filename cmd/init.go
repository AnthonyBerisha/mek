package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create the configuration folder and file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		commandInit()
	},
}

func commandInit() {
	var folderPath = getFolderPath()

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		println("Folder does not exist, creating...", err)

		err := createFolder(folderPath)
		if err != nil {
			log.Fatal("Error creating folder at:", folderPath, err)
		}
	}

	// Now we can create the file
	configFilePath := filepath.Join(folderPath, "mek.yaml")
	var fileExists = checkFileExists(configFilePath)
	if fileExists {
		println("File already exists at path:", configFilePath)
	} else {
		println("File does not exist, creating then write")
		println(configFilePath)

		err := createFile(configFilePath)
		if err != nil {
			log.Fatal("Error creating file:", err)
		} else {
			println("File created successfully")
		}
	}
}

func createFile(path string) error {
	_, err := os.Create(path)
	return err
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

func getFolderPath() string {
	configDir, errConfigDir := os.UserHomeDir()
	if errConfigDir != nil {
		log.Fatal("Error getting config dir:", errConfigDir)
	}

	configPath := filepath.Join(configDir, "/.config/mek")

	return configPath
}

func createFolder(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal("Error creating config folder:", err)
	}

	return err
}

func checkInit() error {
	var err error

	configPath := getFolderPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.New("configuration folder does not exist, please run 'mek init' to create it")
	}

	configFilePath := filepath.Join(configPath, "mek.yaml")
	if !checkFileExists(configFilePath) {
		return errors.New("configuration file does not exist, please run 'mek init' to create it")
	}

	return err
}
