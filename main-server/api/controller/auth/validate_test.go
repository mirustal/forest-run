package auth

import (
	"forest-run/main-server/api/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		r     protocol.SignUpRequest
		valid bool
	}{
		{
			protocol.SignUpRequest{
				Username: "test",
				Password: "test",
			},
			true,
		},
		{
			protocol.SignUpRequest{
				Username: "",
				Password: "test",
			},
			false,
		},
		{
			protocol.SignUpRequest{
				Username: "test",
				Password: "",
			},
			false,
		},
		{
			protocol.SignUpRequest{},
			false,
		},
		{
			protocol.SignUpRequest{
				Username: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
				Password: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
			},
			true,
		},
		{
			protocol.SignUpRequest{
				Username: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
				Password: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
			},
			false,
		},
		{
			protocol.SignUpRequest{
				Username: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
				Password: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
			},
			false,
		},
	}

	for _, c := range cases {
		err := Validate(c.r)
		if c.valid {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}
