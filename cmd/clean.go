/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean all stored cache",
	Long: `The clean command clears all tulip cache stored in the system location. `,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		dir := usr.HomeDir
		dirpath := filepath.Join(dir, ".tulip", "tmp")
		os.RemoveAll(dirpath)
		fmt.Println("Cleared all Cache")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
