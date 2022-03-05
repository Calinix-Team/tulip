/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Resync software repositories",
	Long: `Refresh your package databases to update to the latest fetches of software.`,
	Run: func(cmd *cobra.Command, args []string) {
		cm := exec.Command("pkexec", "pacman", "-Sy")
		cm.Stdout = os.Stdout
		cm.Stdin = os.Stdin
		cm.Stderr = os.Stderr
		err := cm.Run()
		if err!=nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
