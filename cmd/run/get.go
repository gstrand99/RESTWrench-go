/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package run

import (
  "fmt"
  "github.com/spf13/viper"
	"github.com/spf13/cobra"
  "github.com/gstrand99/RESTWrench-go/fuctionality/requests"
)

var urlFlag string
var baseFlag bool

// getCmd represents the get command
var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Makes a GET request",
    Long:  `This command makes a GET request to the given URL.`,
    Run: func(cmd *cobra.Command, args []string) {
        var url string
        if baseFlag {
            baseURL := viper.GetString("baseurl")
            if baseURL == "" {
                fmt.Println("No base URL is set in the config file.")
                return
            }
            url = baseURL + urlFlag
        } else {
            url = urlFlag
        }

        requests.Get("GET", "single request", url)
    },
}
func init() {
	RunCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
  getCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "URL to make a GET request to")
  getCmd.Flags().BoolVarP(&baseFlag, "base", "b", false, "Use base URL from config file")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
