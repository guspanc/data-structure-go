package stack

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		expected interface{}
	}{
		{
			desc:     "",
			expected: "last",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			stack := Stack{}
			stack.Push("foo")
			stack.Push("bar")
			stack.Push(tC.expected)
			last := stack.Pop()
			if last != tC.expected {
				t.Errorf("expected: %s got: %s", tC.expected, last)
			}
		})
	}
}
