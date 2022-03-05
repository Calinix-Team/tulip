/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Calinix-Team/tulip/internal"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an installed package from your system",
	Long: `Removes the specified packages from the system along with any packages depending on the packages being removed.
	
To use pacman flags here (like --needed, --noprogressbar, etc.), follow the synopsis below:
	tulip remove <package> -- <flags>

Example:
	tulip remove alacritty -- --noprogress 

Remember that for remove command specific flags you do not need the "-- <flags>" format, you can just provide the flags without the extra "--"
`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PacmanRemove(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
