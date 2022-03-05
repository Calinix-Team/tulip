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

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your whole system",
	Long: `Upgrade all installed packages in your system and keep your whole system up-to-date
	
Note: DO NOT interrupt the process of updating that may lead to system failure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cm := exec.Command("pkexec", "pacman", "-Syu")
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
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
