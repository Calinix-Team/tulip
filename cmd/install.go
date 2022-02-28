/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
)

import "github.com/Calinix-Team/tulip/internal"

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package from local repositories",
	Long: `The install command is used to download the latest version of your desired application 
from an online software repository pointed to by your /etc/pacman.d configuration file and
install that application on your Linux machine.

To use flags here, follow the synopsis below:
	tulip install <package> -- <flags>
`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PacmanInstall(args)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.PersistentFlags().Bool("needed", false, "do not reinstall up to date packages")
	installCmd.PersistentFlags().String("dbpath", "", "set an alternate database location")
	installCmd.PersistentFlags().Bool("clean", false, "remove old packages from cache directory (-cc for all)")
	installCmd.PersistentFlags().Bool("nodeps", false, "skip dependency version checks (-dd to skip all checks)")
	installCmd.PersistentFlags().Bool("groups", false, "view all members of a package group")
	installCmd.PersistentFlags().Bool("arch", false, "set an alternate architecture")
	installCmd.PersistentFlags().Bool("noprogressbar", false, "do not show a progress bar when downloading files")
	installCmd.PersistentFlags().Bool("noscriptlet", false, "do not execute the install scriptlet if one exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
