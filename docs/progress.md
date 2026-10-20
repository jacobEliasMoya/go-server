# Learning progress

## Current checkpoint: project organization

- The module is named `go-server`; the learning project is Go Server Lab.
- Application code remains in one `main.go` file.
- `main` is empty. There are no routes, handlers, or server process yet.
- Removed the unused `fmt` import and the commented-out `net/http` import from the starter file.
- Changed the Go requirement from 1.27.1 to 1.26.4 to match the installed toolchain. No language features or dependencies require the higher version in this skeleton.
- Added a README and this checkpoint record.
- Kept the existing teaching instructions and ignore rules.

## Verification

Verified using the installed Go 1.26.4 toolchain with automatic toolchain downloading disabled:

- `gofmt -w main.go` completed successfully.
- `go run .` exited successfully with no output.
- `go test ./...` passed compilation and reported `[no test files]`.
- `go vet ./...` completed successfully with no findings.

There is no server behavior to test yet. Git status inspection returned `this operation must be run in a work tree` in this execution environment, despite the repository configuration specifying `bare = false`. No Git configuration or staging changes were made.

## Next checkpoint

Explain and build the first HTTP server on port 8080, with a root response of `Hello from Go`. Wait for the learner to request this checkpoint before adding it.

Understanding goals for that next checkpoint:

- Explain how `package main` and `func main()` identify an executable and its entry point.
- Describe how a request reaches a handler and how the handler writes a response.
- Start the server, request the root URL, and stop the process.

## Scope

Continue the server learning project independently of the separate GitGo CLI project. No framework, database, or additional application packages are needed at this checkpoint.
