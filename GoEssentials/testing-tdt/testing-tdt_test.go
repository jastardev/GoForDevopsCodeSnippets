package main

import (
	"fmt"
	"testing"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name was empty")
	}
	return "Hello " + name, nil
}

func TestGreet(t *testing.T) {
	tests := []struct {
		desc      string //What we are testing
		name      string //The name we will pass
		want      string //What we expect to be returned
		expectErr bool   ///Whether or not we expect an error
	}{
		{
			desc:      "Error: name is an empty string",
			expectErr: true,
			//name and want are "", nil value for strings
		},
		{
			desc: "Success",
			name: "John",
			want: "Hello John",
			//expectErr is set to the zero value, false
		},
	}

	for _, test := range tests {
		got, err := Greet(test.name)
		switch {
		//Did not get an error but expected one
		case err == nil && test.expectErr:
			t.Errorf("TestGreet(%s): g0t err == nil, wanted err != nil", test.desc)
			continue
		//Got an error but did not expect one
		case err != nil && !test.expectErr:
			t.Errorf("TestGreet(%s): got err == %s, wanted err == nil", test.desc, err)
			continue
		//Got an error that we expected
		case err != nil:
			continue
		}

		if got != test.want {
			t.Errorf("TestGreet(%s): got result %q, want %q", test.desc, got, test.want)
		}

	}
}
