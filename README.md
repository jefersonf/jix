# Jira Issue eXtractor 

WIP

```
Usage
  jix [flags]

Flags:
  -f, --format string        Output format (only JSONL and CSV are available) (default "jsonl")
  -h, --help                 help for jix
  -o, --output-path string   Path to the output file (default "./data")
  -p, --project-key string   Jira project key
  -v, --verbose              Set verbose mode

```

### Usage Example

From source, simply clone this repository.

```
cd jix
go build .
```
and then run it.

```
./jix -p TESTGEN -o ./sample -v
```