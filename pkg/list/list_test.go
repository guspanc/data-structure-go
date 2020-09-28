package list

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "default",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			list := &LinkedList{}
			list.Prepend(NewNode(1))
			list.Prepend(NewNode(2))
			list.Prepend(NewNode(3))
			list.DeleteWithValue(2)
			if list.Len() != 2 {
				t.Errorf("expected 2 got %d", list.Len())
			}
		})
	}
}
