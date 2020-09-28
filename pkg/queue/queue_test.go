package queue

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		expected interface{}
	}{
		{
			desc:     "",
			expected: "first",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			queue := Queue{}
			queue.Enqueue(tC.expected)
			queue.Enqueue("foo")
			queue.Enqueue("bar")
			first := queue.Dequeue()
			if first != tC.expected {
				t.Errorf("expected: %s got: %s", tC.expected, first)
			}
		})
	}
}
