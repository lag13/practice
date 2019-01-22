package practice

func StockProfit(stocks []int) (int, int) {
	bestBuy := 0
	bestSell := 1
	bestSoFar := stocks[bestSell] - stocks[bestBuy]
	for buy := 1; buy < len(stocks)-1; buy++ {
		for sell := buy + 1; sell < len(stocks); sell++ {
			profit := stocks[sell] - stocks[buy]
			if profit > bestSoFar {
				bestSoFar = profit
				bestBuy = buy
				bestSell = sell
			}
		}
	}
	return stocks[bestBuy], stocks[bestSell]
}
