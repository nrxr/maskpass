//
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

	msk.User = url.UserPassword(msk.User.Username(), MASK)
	return msk, nil
}

// String returns the url.URL.String() value and an error if something went
// wrong calling Mask.
func String(uri string) (string, error) {
	msk, err := Mask(uri)
	if err != nil {
		return "", err
	}
	return msk.String(), nil
}

// StringNE returns a masked url.URL.String() without error handling. It's a
// wrapper of String().
func StringNE(uri string) string {
	msk, _ := String(uri)
	return msk
}
