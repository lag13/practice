package practice_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lag13/practice"
)

func TestMoveZerosLeft(t *testing.T) {
	tests := []struct {
		arr     []int
		wantArr []int
	}{
		{
			arr:     []int{0, 1, 0, 2, 0, 3},
			wantArr: []int{0, 0, 0, 1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("MoveZerosLeft(%+v)", test.arr), func(t *testing.T) {
			practice.MoveZerosLeft(test.arr)
			if got, want := test.arr, test.wantArr; !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}
