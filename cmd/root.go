package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
    Use:   "mainflux-cli",
    Short: "CLI for Mainflux",
    Long: `Mainflux Command Line Interface`,
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
    },
}
