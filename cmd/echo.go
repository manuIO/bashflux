package cmd

import (
        "github.com/spf13/cobra"
        "fmt"
        "strings"
)

var EchoCmd = &cobra.Command{
    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Long:  `echo is for echoing anything back.
    Echo echo’s.
    `,
    Run: echoRun,
}

func echoRun(cmd *cobra.Command, args []string) {
    fmt.Println(strings.Join(args, " "))
}

func init() {
    RootCmd.AddCommand(EchoCmd)
}
