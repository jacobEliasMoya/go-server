# Go Server Lab

A small learning project for building Go backend skills from first principles, using the standard library and one checkpoint at a time.

## Current state

The project contains a minimal executable skeleton. Running it exits without output. It does not start an HTTP server yet.

## Files

```text
go-server/
  main.go           Program entry point
  go.mod            Module name and minimum Go version
  .gitignore        Files excluded from Git
  AGENTS.md         Teaching and project instructions
  README.md         Setup and current behavior
  docs/
    progress.md     Learning checkpoint and next step
```

Keep application code in `main.go` while learning the first server concepts. Add packages only when the working code needs them.

## Run

In PowerShell:

```powershell
cd C:\Users\Jemoy\Golang\go-server
go version
go run .
```

The module requires Go 1.26.4 or later. The toolchain found during setup was Go 1.26.4 on Windows/amd64.

Expected behavior today: the program exits successfully with no output.

## Check

```powershell
go test ./...
go vet ./...
```

There are no tests yet. `go test` currently checks that the package compiles; it does not verify server behavior.

## Next learning checkpoint

Build the first HTTP server in `main.go`, listen on port 8080, and return `Hello from Go` for the root route. Explain packages, imports, `func main`, handler functions, and starting the server as part of that checkpoint.

See [progress notes](docs/progress.md) for the current stopping point and [AGENTS.md](AGENTS.md) for the teaching approach.
