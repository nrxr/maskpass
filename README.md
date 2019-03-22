# maskpass

Simple usage:

    package main

    import (
      "fmt"
      "github.com/nrxr/maskpass"
    )

    func main() {
      uri := "https://user:password@localhost:port/path"
      fmt.Println(maskpass.String(uri))
      // Output: "https://user:xxxxxx@localhost:port/path"
    }

Full docs on https://godoc.org/github.com/nrxr/maskpass
