package leetcode_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestAddTwoNumbers(t *testing.T) {
	tt := []struct {
		l1   *leetcode.ListNode
		l2   *leetcode.ListNode
		want *leetcode.ListNode
	}{
		{
			l1:   leetcode.NumToListNode(342),
			l2:   leetcode.NumToListNode(465),
			want: leetcode.NumToListNode(807),
		},
		{
			l1:   leetcode.NumToListNode(0),
			l2:   leetcode.NumToListNode(2489),
			want: leetcode.NumToListNode(2489),
		},
		{
			l1:   leetcode.NumToListNode(99),
			l2:   leetcode.NumToListNode(99),
			want: leetcode.NumToListNode(198),
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("AddTwoNumbers(%s, %s)", tc.l1, tc.l2), func(t *testing.T) {
			if got, want := leetcode.AddTwoNumbers(tc.l1, tc.l2), tc.want; !reflect.DeepEqual(got, want) {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
