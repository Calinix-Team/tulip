/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	// "os"
	"github.com/Calinix-Team/tulip/internal"
	"github.com/fatih/color"
	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build and Install a Package from the User Repositories (AUR)",
	Long: `Build packages from AUR and install them with their dependencies.

The Arch User Repository (AUR) is a user maintained repository with a lot of external and development versions of packages which are not present in repositories specified by pacman. This build command fetches packages from the AUR and installs them in the local system.`,
	Run: func(cmd *cobra.Command, args []string) {
		bd, _ := cmd.Flags().GetString("builddir")
		keep, _ := cmd.Flags().GetBool("keep")
		res, _ := requests.Get(fmt.Sprintf("https://aur.archlinux.org/packages/%v", args[0]))
		if res.StatusCode == 404 {
			color.Red("Could Not Find Package: %v", args[0])
		} else {
			for i := 0; i < len(args); i++ {
				internal.BuildFromAUR(args[i], bd, keep)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	buildCmd.Flags().StringP("builddir", "b", "", "specify the directory the package should be built in")
	buildCmd.Flags().BoolP("keep", "k", false, "save build files in cache/build directory")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
