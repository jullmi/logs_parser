package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
    Use:   "stats",
    Short: "Show log statistics for a given date",
    Run: func(cmd *cobra.Command, args []string) {
        if date == "" {
            fmt.Println("Please provide a date using --date (format: YYYY-MM-DD)")
            return
        }

        file, err := os.Open(filePath)
        if err != nil {
            fmt.Printf("Failed to open log file: %v\n", err)
            return
        }
        defer file.Close()

        // Counters for different log levels
        var infoCount, errorCount, warnCount int

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            line := scanner.Text()
            if strings.HasPrefix(line, date) {
                if strings.Contains(line, "INFO") {
                    infoCount++
                } else if strings.Contains(line, "ERROR") {
                    errorCount++
                } else if strings.Contains(line, "WARN") {
                    warnCount++
                }
            }
        }

        if err := scanner.Err(); err != nil {
            fmt.Printf("Error reading file: %v\n", err)
        }

        // Output stats
        fmt.Printf("Log stats for %s:\n", date)
        fmt.Printf("INFO: %d\n", infoCount)
        fmt.Printf("ERROR: %d\n", errorCount)
        fmt.Printf("WARN: %d\n", warnCount)
    },
}

func init() {
    rootCmd.AddCommand(statsCmd)

    // Register flags for stats command
    statsCmd.Flags().StringVarP(&filePath, "file", "f", "app.log", "Path to log file")
    statsCmd.Flags().StringVarP(&date, "date", "d", "", "Date to filter logs (YYYY-MM-DD)")
}