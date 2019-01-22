package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

func TestStockProfit(t *testing.T) {
	tests := []struct {
		nums     []int
		wantBuy  int
		wantSell int
	}{
		{
			nums:     []int{8, 5, 12, 9, 19, 1},
			wantBuy:  5,
			wantSell: 19,
		},
		{
			nums:     []int{21, 12, 11, 9, 6, 3},
			wantBuy:  12,
			wantSell: 11,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("StockProfit(%+v)", test.nums), func(t *testing.T) {
			buy, sell := practice.StockProfit(test.nums)
			if got, want := buy, test.wantBuy; got != want {
				t.Errorf("want buy %d, got %d", got, want)
			}
			if got, want := sell, test.wantSell; got != want {
				t.Errorf("want sell %d, got %d", got, want)
			}
		})
	}
}
