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

## Example

Create a `main.go` file with the following contents:

```go
package main

import (
  "log"

  "github.com/patrickjmcd/go-version"
)

func main() {
  log.Printf(
    "Starting the test...\n%s",
    version.GetVersion(),
  )
}
```

Build the executable with the following command:

```shell
go build \
  -ldflags "-s -w -X github.com/patrickjmcd/go-version/version.Release=0.0.1 \
  -X github.com/patrickjmcd/go-version/version.Commit=aabbddcceeff -X github.com/patrickjmcd/go-version/version.BuildTime=2021-11-09_21:18:54 \
  -o testversion
```

Run the executable and see the output!

```shell
$ ./testversion
Starting the test...
commit: aabbddcceeff, build time: 2021-11-09_21:18:54, release: 0.0.1
```
