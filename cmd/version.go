package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var version = "v0.1.0"

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version of logs_parser",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("logs_parser", version)
    },
}


func init() {
    rootCmd.AddCommand(versionCmd)
}