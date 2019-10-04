# maskpass

This package masks passwords from URL strings making them log-friendly.

Even if you're using `url.URL` structures instead of plain strings, this is very
useful, you just need to pass the `url.URL.String()` value and `maskpass` will
take care of it.

It will only mask the URL if there's a password in the URL, otherwise it will
just output the same thing `url.URL.String()` would.

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
