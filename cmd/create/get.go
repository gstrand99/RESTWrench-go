/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package create

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "io/ioutil"
    "os"
)

type GetRequest struct {
    REQUESTTYPE string `json:"type"`
    NAME string `json:"name"`
    URL string `json:"url"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Create a new GET request",
    Long:  `This command creates a new GET request and saves it to a JSON file.`,
    Args:  cobra.ExactArgs(3), // Expect exactly one argument
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        requestType := args[1]
        url := args[2]
          
        dir := "wrench/requests"
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            err := os.MkdirAll(dir, 0755)
            if err != nil {
              fmt.Printf("Error creating directory: %v\n", err)
              return
            }
         }
    

        var getRequests []GetRequest
        filename := "wrench/requests/get_requests.json"

        // Check if the file exists
        if _, err := os.Stat(filename); err == nil {
            // The file exists, so read it and unmarshal the JSON
            data, err := ioutil.ReadFile(filename)
            if err != nil {
                fmt.Printf("Error reading file: %v\n", err)
                return
            }

            err = json.Unmarshal(data, &getRequests)
            if err != nil {
                fmt.Printf("Error unmarshalling JSON: %v\n", err)
                return
            }
        }

        // Add the new GET request
        getRequests = append(getRequests, GetRequest{REQUESTTYPE: requestType, NAME: name, URL: url})

        // Marshal the GET requests to JSON
        data, err := json.MarshalIndent(getRequests, "", "  ")
        if err != nil {
            fmt.Printf("Error marshalling JSON: %v\n", err)
            return
        }

        // Write the JSON to the file
        err = ioutil.WriteFile(filename, data, 0644)
        if err != nil {
            fmt.Printf("Error writing file: %v\n", err)
            return
        }

        fmt.Printf("Added GET request to %s\n", url)
    },
}

func init() {
    CreateCmd.AddCommand(getCmd)
}
