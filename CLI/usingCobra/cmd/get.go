/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GoForDevOps/RPC/gRPC/client"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		const devAddr = "127.0.0.1:3450"

		fs := cmd.Flags()
		addr := mustString(fs, "addr")

		if mustBool(fs, "dev") {
			addr = devAddr
		}

		c, err := client.New(addr)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		a, q, err := c.QOTD(cmd.Context(), mustString(fs, "author"))
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		switch {
		case mustBool(fs, "json"):
			b, err := json.Marshal(
				struct {
					Author string
					Quote  string
				}{a, q})
			if err != nil {
				panic(err)
			}
			fmt.Println("%s\n", b)
		default:
			fmt.Println("Author: ", a)
			fmt.Println("Quote: ", q)
		}
	},
}

func mustString(fs *pflag.FlagSet, name string) string {
	v, err := fs.GetString(name)
	if err != nil {
		panic(err)
	}
	return v
}

func mustBool(fs *pflag.FlagSet, name string) bool {
	v, err := fs.GetBool(name)
	if err != nil {
		panic(err)
	}
	return v
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("dev", "d", false, "Usees the dev server instead of prod")
	getCmd.Flags().String("addr", "127.0.0.1:80", "Sets the QOTD Server to use, defaults to production")
	getCmd.Flags().StringP("author", "a", "", "Specify the author to get a quote for")
	getCmd.Flags().Bool("json", false, "Output is in JSON format")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
