package maskpass

import (
	"testing"
)

func TestMask(t *testing.T) {
	var data = []struct {
		uri  string
		user string
	}{
		{uri: "https://asd:ius@hola.com/jjdk", user: "asd"},
		{uri: "amqp://awd:uuu@hey.com/", user: "awd"},
	}

	for _, d := range data {
		msk, _ := Mask(d.uri)
		if msk.User.Username() != d.user {
			t.Errorf("Mask(%s) changed User value: expected %s and got %s", d.uri, d.user, msk.User.Username())
		}

		p, has := msk.User.Password()
		if !has {
			t.Errorf("getting Password of %s: %t", d.uri, has)
		}
		if p != MASK {
			t.Errorf("Mask(%s) did not mask the password: expected %s and got %s", d.uri, MASK, p)
		}
	}
}

var dataString = []struct {
	uri string
	exp string
}{
	{
		uri: "mysql://user:password@host.tld/database",
		exp: "mysql://user:xxxxxx@host.tld/database",
	},
	{ // if there's no password then it should not mask anything
		uri: "mysql://user@host.tld/database",
		exp: "mysql://user@host.tld/database",
	},
}

func TestStringError(t *testing.T) {
	for _, d := range dataString {
		msk, err := StringError(d.uri)
		if err != nil {
			t.Errorf(`String("%s") returned an error: %v`, d.uri, err)
		}

		if msk != d.exp {
			t.Errorf(`String("%s") returned wrong result. Got %s and expected %s`, d.uri, msk, d.exp)
		}
	}
}

func TestString(t *testing.T) {
	for _, d := range dataString {
		msk := String(d.uri)
		if msk != d.exp {
			t.Errorf(`String("%s") returned wrong result. Got %s and expected %s`, d.uri, msk, d.exp)
		}
	}
}
