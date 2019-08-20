# maskpass

`maskpass` will mask a given string with `xxxxxx`, this is specially useful to
hide passwords on logs.

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

## License

ⓒ 2019, nrxr `<nrxr@disroot.org>`. Released under the MIT license terms.
