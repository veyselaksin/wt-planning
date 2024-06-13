/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mocky-cli/helpers/requests"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch command is used to fetch mock data from an URL and save it to a database.",
	Long: `
		Fetch command is used to fetch mock data from an URL and save it to a database.
		Default Database is Postgres.
`,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			return
		}
		save, err := cmd.Flags().GetBool("save")
		if err != nil {
			return
		}
		fetch(url, save)
	},
}

func fetch(url string, save bool) {
	// Fetch the data from the URL
	res, err := requests.NewRequest(requests.Request{
		URL:    url,
		Method: "GET",
		Body:   nil,
		Header: nil,
	})
	if err != nil {
		return
	}

	fmt.Println(string(res))
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.
	fetchCmd.Flags().StringP("url", "u", "", "URL to fetch mock data from")
	fetchCmd.Flags().BoolP("save", "s", false, "Save the fetched data to a database")
	err := fetchCmd.MarkFlagRequired(`url`)
	if err != nil {
		return
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
