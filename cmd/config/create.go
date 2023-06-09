/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a configuration file",
    Long:  `This command creates a configuration file with some default values.`,
    Run: func(cmd *cobra.Command, args []string) {
        filename := "wrench/config.yaml"

        // Check if the file already exists
        if _, err := os.Stat(filename); err == nil {
            fmt.Println("Config file already exists.")
            return
        }

        // Set default values
        viper.Set("baseurl", "https://api.example.com")

        // Create the directory if it doesn't exist
        dir := "wrench/"
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            err = os.MkdirAll(dir, 0755)
            if err != nil {
                fmt.Printf("Error creating directory: %v\n", err)
                return
            }
        }

        // Write the configuration to a file
        err := viper.WriteConfigAs(filename)
        if err != nil {
            fmt.Printf("Error writing config file: %v\n", err)
            return
        }

        fmt.Println("Config file created.")
    },
}

func init() {
	// Here you will define your flags and configuration settings.
  ConfigCmd.AddCommand(createCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
