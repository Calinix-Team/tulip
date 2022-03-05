/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/Calinix-Team/tulip/internal"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search packages from either the official or user repositories",
	Long: `Search package metadata for keywords. Keywords are matched as case-insensitive substrings, globbing is supported. `,
	Run: func(cmd *cobra.Command, args []string) {
		aurs, _ := cmd.Flags().GetBool("aur")

		if !aurs {
			arg := []string{"pacman", "-Ss"}
			argsl := append(arg, args...)
			cm := exec.Command("pkexec", argsl...)
			cm.Stdout = os.Stdout
			cm.Stderr = os.Stderr
			cm.Stdin = os.Stdin
			cm.Run()
		} else {
			internal.GetResults("desc", "https://aur.archlinux.org", true, false, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().BoolP("aur", "a", false, "Search in external User Repositories (AUR)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
