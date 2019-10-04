// Package maskpass masks the password of an URL, so something like
// https://user:password@host:port/ would end up like
// https://user:xxxxxx@host:port/.
//
// Usage:
//
// The easiest way to use maskpass is with the method String.
//
//  uri := url.Parse("whatever")
//  masked := maskpass.String(uri.String())
package maskpass

import (
	"net/url"
)

const MASK = "xxxxxx"

// Mask returns an *url.URL with the password masked to the placeholder
// constant's value.
func Mask(uri string) (*url.URL, error) {
	msk, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if _, has := msk.User.Password(); has {
		msk.User = url.UserPassword(msk.User.Username(), MASK)
	}
	return msk, nil
}

// StringError returns the url.URL.String() value and an error if something
// went wrong calling Mask.
func StringError(uri string) (string, error) {
	msk, err := Mask(uri)
	if err != nil {
		return "", err
	}
	return msk.String(), nil
}

// String returns a masked url.URL.String() without error handling. It's a
// wrapper of String().
func String(uri string) string {
	msk, _ := StringError(uri)
	return msk
}
