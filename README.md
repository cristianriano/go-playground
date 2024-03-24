# Go Playground

Simple project with examples and experiments around the Go language.

## Setup

1. Check [.tool-versions](.tool-versions) for the Golang used version.
2. Run main with `go run .` or build an executable with `go build` (dependencies are downloaded automatically)
3. To install dependencies declared in **go.mod** explicitly use `go mod download`.

### How to Modules
To start a new project in Go

1. Create a new module `go mod init <module>`
   1. Module name by conventions is `<host>/<user>/<module>`
   2. Where host is github.com, user or company that manages the module and finally module name.
   3. This will create a **go.mod** file with the module name and the Go version used. It will also contain dependencies and their versions
2. Every module needs a **main package** and a **main function** which will be the entry point (file doesn't need to be called main.go)
3. To add a dependency, use `go get <module>` it will be added to the **go.mod** file and the **go.sum** file (which contains the checksum for integrity check)
4. Only methods and variables **starting with upper case can be used from other modules**

Other useful commands:
1. `go mod tidy` to remove unused dependencies or add missing ones.
2. `go mod vendor` to create a vendor folder with all the dependencies.
3. `go install` downloads and install a package. It will be available in **$GOPATH/bin**
4. `gofmt -w .` format