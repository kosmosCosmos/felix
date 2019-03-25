package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// httpInfoCmd represents the httpInfo command
var httpInfoCmd = &cobra.Command{
	Use:   "httpInfo",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("httpInfo called")
	},
}

func init() {
	rootCmd.AddCommand(httpInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
