
# logы_parser

A simple CLI tool to filter logs by date and level, built with Go and Cobra.

## 💻 Usage

```bash
go run main.go --file=app.log --date=2025-04-10 --level=ERROR --out=filtered.log


🛠 Flags
	•	--file / -f: path to the log file (default: app.log)
	•	--date / -d: date to filter logs (YYYY-MM-DD)
	•	--level / -l: log level (INFO, ERROR, etc.)
	•	--out / -o: output file path
    