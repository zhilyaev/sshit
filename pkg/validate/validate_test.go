package validate

import "testing"

type pairIsDomain struct {
	input  string
	output bool
}

var testsIsDomain = []pairIsDomain{
	{"dmitriy.zhilyaev.ru", true},
	{"abc.xyz", true},
	{"1234", false},
	{"", false},
}

func TestIsDomain(t *testing.T) {
	for _, test := range testsIsDomain {
		if IsDomain(test.input) != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", IsDomain(test.input),
			)
		}
	}
}
