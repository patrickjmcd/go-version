# Go Version Package Printer

A simple package to show commit, version, and build time for a go program

## Usage

In your program, include `github.com/patrickjmcd/go-version`.

In your `main.go` file's `main` function, include

```go
fmt.Println(version.GetVersion())
```

When you build your program, build using a command like:

```shell
go build \
  -ldflags "-s -w -X github.com/patrickjmcd/go-version/version.Release=${RELEASE} \
  -X github.com/patrickjmcd/go-version/version.Commit=${COMMIT} -X github.com/patrickjmcd/go-version/version.BuildTime=${BUILD_TIME}" \
  -o ${APP}

```
