/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package run

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "io/ioutil"
    "github.com/gstrand99/RESTWrench-go/fuctionality/requests"
)

type GetRequest struct {
    Type string `json:"type"`
    Name string `json:"name"`
    URL string `json:"url"`
}

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {
        filename := "wrench/requests/get_requests.json"

        // Read the file
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Printf("Error reading file: %v\n", err)
            return
        }

        // Unmarshal the JSON
        var getRequests []GetRequest
        err = json.Unmarshal(data, &getRequests)
        if err != nil {
            fmt.Printf("Error unmarshalling JSON: %v\n", err)
            return
        }

        // Run all GET requests
        for _, getRequest := range getRequests {
            requests.Get(getRequest.Type, getRequest.Name, getRequest.URL)
        }
    },
}
func init() {
	RunCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
