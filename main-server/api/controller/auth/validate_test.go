package auth

import (
	"github.com/stretchr/testify/assert"
	"main-server/domain"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		r     domain.SignUpRequest
		valid bool
	}{
		{
			domain.SignUpRequest{
				Username: "test",
				Password: "test",
			},
			true,
		},
		{
			domain.SignUpRequest{
				Username: "",
				Password: "test",
			},
			false,
		},
		{
			domain.SignUpRequest{
				Username: "test",
				Password: "",
			},
			false,
		},
		{
			domain.SignUpRequest{},
			false,
		},
		{
			domain.SignUpRequest{
				Username: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
				Password: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
			},
			true,
		},
		{
			domain.SignUpRequest{
				Username: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
				Password: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
			},
			false,
		},
		{
			domain.SignUpRequest{
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
