/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"log"
)

import "github.com/Calinix-Team/tulip/internal"

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package from local repositories",
	Long: `The install command is used to download the latest version of your desired application from an online software repository pointed to by your /etc/pacman.d configuration file and install that application on your Linux machine.

To use pacman flags here (like --needed, --noprogressbar, etc.), follow the synopsis below:
	tulip install <package> -- <flags>

Example:
	tulip install alacritty -- --noprogress 

Remember that for install command specific flags you do not need the "-- <flags>" format, you can just provide the flags without the extra "--"
`,
	Run: func(cmd *cobra.Command, args []string) {
		upst, _ := cmd.Flags().GetBool("nosync")
		if !upst{
		cm := exec.Command("pkexec", "pacman", "-Sy")

		cm.Stdout = os.Stdout
		cm.Stdin = os.Stdin
		cm.Stderr = os.Stderr
		err := cm.Run()
		if err!=nil {
			log.Fatal(err)
		}
	}
		internal.PacmanInstall(args)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.PersistentFlags().Bool("nosync", false, "do not sync repositories before installation")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}