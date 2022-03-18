# Day 0 - Setup

## Installing golang

- To write go code we need to setup a go environment, with go development tools installed.
- You can usually install the latest from [golang.org/dl](https://golang.org/dl/).
- If you are on a mac, you can install golang from the [homebrew](https://brew.sh/) package manager.
     - `brew install go`
- Check the installation of go by running:
     - `go version`

## Go Workspace
- Go went through a lot of iterations about what it means to be a workspace.
- It is important to understand that a workspace is a directory that contains a go project.
- You are free to organize your go projects in any place you want.
- Go still expects there to be a **single** workspace for third-party Go tools installed via `go install`
- By default the workspace is located at `$HOME/go`.
     - Source code stored in `$HOME/go/src`
     - Binaries stored in `$HOME/go/bin`
- You can change the default by setting `$GOPATH` environment variable.
- Some online resources tell you to set the `$GOROOT` environment variable. This variable specifies the location where your Go development environment is installed. This is no longer necessary; the `go` tool figures this out automatically.

## The `go` Command
- Go ships with many developmental tools.
- The `go` command is the entry point to all of these tools.
- They include a compiler, code formatter, linter, dependency manager, test runner, and more.
### `go run` and `go build`
- `go run` and `go build` take in either single file, list of files or the name of the package.
- `go run` runs the main function of the package.
- `go build` compiles the package and builds a binary.
- `go run` also compiles the code into a binary. However, the binary is stored in a temporary directory.
- The `go run` command builds the binary, executes the binary from that temporary directory, and then deletes the binary after your program finishes. 
- This makes the `go run` command useful for testing out small programs or using Go like a scripting language.
- When you use `go build`, the name of the **binary matches the name of the file or package** you passed in.
- You can change that buy passing in a `-o` flag.
     - `go build -o hello_world hello.go` will create a binary named `hello_world`
- You can also pass in a `-i` flag to `go build` to install the binary into your `$GOPATH/bin` directory.
- You can also pass in a `-v` flag to `go build` to get more information about the compilation process.
- Use `go build` to create a binary that is distributed for other people to use.

### Third party go tools
- You can use third party go tools by using the `go install` command.
- `go install` takes in the name of the package.
- `go install` will install the package into your `$GOPATH/bin` directory.
- Go’s method for publishing code is a bit different than most other languages. Go developers don’t rely on a centrally hosted service, like Maven Central for Java or the NPM registry for JavaScript. Instead, they share projects via their source code repositories.
- The `go install` command takes an argument, which is the location of the source code repository of the project you want to install, followed by an `@` and the version of the tool you want (if you just want to get the latest version, use `@latest`). It then downloads, compiles, and installs the tool into your `$GOPATH/bin` directory.

### Formatting code
- You can use the `go fmt` command to format your code.
- It automatically formats your code to match the standard format.
- There is a enhanced version of `go fmt` called `goimports`.
- `goimports` automatically imports your code into the standard format. Removes unused imports, arranges in alphabetical order, and attempts to guess any unspecified imports.
- Install `go imports` by running:
     - `go install golang.org/x/tools/cmd/goimports@latest`
- Run it `goimports -l -w .` to format your code.
- The `-l` flag tells goimports to print the files with incorrect formatting to the console. The `-w` flag tells goimports to modify the files in-place. The `.` specifies the files to be scanned: everything in the current directory and all of its subdirectories.

### Semicolon insertion rule
- Like C or Java, Go requires a semicolon at the end of every statement. 
- However, Go developers never put the semicolons in themselves; the Go compiler does it for them following a very simple rule described in [Effective Go](https://go.dev/doc/effective_go#semicolons).
- If the last token before a newline is any of the following, the lexer inserts a semicolon after the token:
     - An identifier (which includes words like `int` and `float64`)
     - A basic literal such as a number or string constant
     - One of the tokens: `break`, `continue`, `fallthrough`,`return`, `++`, `--`, `)`, or `}`

### Linting and Vetting
- `go fmt` and `go lint` are two of the most important tools in the Go ecosystem.
- `go fmt` is used to format your code.
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- `golint` tries to ensure your code is well-formed and follows the style guidelines
- `golint ./...` will lint the current directory and all of its subdirectories.
- `go vet` is a tool that checks for common mistakes in Go code. Includes things like passing the wrong number of parameters to formatting methods or assigning values to variables that are never used.
- `golangci-lint` combines `golint` and `go vet` into a single tool and an ever-increasing set of other code quality tools
- 