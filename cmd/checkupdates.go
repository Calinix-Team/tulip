/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/exec"
	"log"

	"github.com/spf13/cobra"
)

// checkupdatesCmd represents the checkupdates command
var checkupdatesCmd = &cobra.Command{
	Use:   "checkupdates",
	Short: "List all packages that need to be upgraded",
	Long: `Safely check all updatable packages without installing them into your system`,
	Run: func(cmd *cobra.Command, args []string) {
		cm := exec.Command("pkexec", "pacman", "-Qu")
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
	rootCmd.AddCommand(checkupdatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkupdatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkupdatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
