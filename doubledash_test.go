package doubledash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitArg(t *testing.T) {

	assert := assert.New(t)

	var tests = []struct {
		in   []string
		real []string
		xtra []string
	}{
		{
			[]string{"-a", "-b"},
			[]string{"-a", "-b"},
			[]string{},
		},
		{
			[]string{"-a", "-b", "--"},
			[]string{"-a", "-b"},
			[]string{},
		},
		{
			[]string{"-a", "-b", "--", "-c"},
			[]string{"-a", "-b"},
			[]string{"-c"},
		},
		{
			[]string{"-a", "-b", "--", "-c", "--", "-d", "-e"},
			[]string{"-a", "-b"},
			[]string{"-c", "--", "-d", "-e"},
		},
	}

	for _, test := range tests {

		real, xtra := Split(test.in)

		msg := fmt.Sprintf("%v", test.in)

		assert.Equal(test.real, real, msg)
		assert.Equal(test.xtra, xtra, msg)

	}

}
