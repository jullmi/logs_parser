package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

var (
    filePath string
    logLevel string
    date     string
    output   string
)

var rootCmd = &cobra.Command{
    Use:   "logs_parser",
    Short: "Filter logs by date and level",
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

        outFile, err := os.Create(output)
        if err != nil {
            fmt.Printf("Failed to create output file: %v\n", err)
            return
        }
        defer outFile.Close()

        writer := bufio.NewWriter(outFile)
        scanner := bufio.NewScanner(file)
        count := 0

        for scanner.Scan() {
            line := scanner.Text()
            if strings.Contains(line, logLevel) && strings.HasPrefix(line, date) {
                writer.WriteString(line + "\n")
                count++
            }
        }

        writer.Flush()

        if err := scanner.Err(); err != nil {
            fmt.Printf("Error reading file: %v\n", err)
        }

        fmt.Printf("Done! Found %d log entries written to %s\n", count, output)
    },
}


func init() {
    rootCmd.Flags().StringVarP(&filePath, "file", "f", "app.log", "Path to log file")
    rootCmd.Flags().StringVarP(&logLevel, "level", "l", "ERROR", "Log level to filter (INFO, ERROR, etc)")
    rootCmd.Flags().StringVarP(&date, "date", "d", "", "Date to filter logs (format: YYYY-MM-DD)")
    rootCmd.Flags().StringVarP(&output, "out", "o", "filtered.log", "Output file to write results")
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

