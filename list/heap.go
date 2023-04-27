package list

type CostProfit struct {
	cost, profit int
}

func MaxMoney(cost, profit []int, money, choice int) int {
	list := make([]*CostProfit, 0)
	for i, value := range cost {
		list = append(list, &CostProfit{
			cost:   value,
			profit: profit[i],
		})
	}

	for choice > 0 && money > 0 && len(list) > 0 {
		initCostHeap(list)
		profitList := getCostHeap(list, money)
		initProfitHeap(profitList)

		item := profitList[0]
		money += item.profit
		choice--

		tmp := make([]*CostProfit, 0)
		for _, value := range list {
			if item != value {
				tmp = append(tmp, value)
			}
		}

		list = tmp
	}

	return money
}

func initCostHeap(list []*CostProfit) {
	begin := len(list)/2 - 1
	for i := begin; i >= 0; i-- {
		adjustCostHeap(list, i, len(list))
	}
}

func adjustCostHeap(list []*CostProfit, index, limit int) {
	left, right := index*2+1, index*2+2
	target := index

	if left < limit && list[left].cost < list[target].cost {
		target = left
	}

	if right < limit && list[right].cost < list[target].cost {
		target = right
	}

	if target != index {
		list[target], list[index] = list[index], list[target]
		adjustCostHeap(list, target, limit)
	}
}

func getCostHeap(list []*CostProfit, money int) []*CostProfit {
	ans := make([]*CostProfit, 0)
	size := len(list)

	for i := 0; i < size; i++ {
		item := list[0]
		if item.cost > money {
			break
		}

		ans = append(ans, list[0])
		list[0] = list[size-1]
		size--
		adjustCostHeap(list, 0, size)
	}

	return ans
}

func initProfitHeap(list []*CostProfit) {
	begin := len(list)/2 - 1
	for i := begin; i >= 0; i-- {
		adjustProfitHeap(list, i, len(list))
	}
}

func adjustProfitHeap(list []*CostProfit, index, limit int) {
	left, right := index*2+1, index*2+2
	target := index

	if left < limit && list[left].profit > list[target].profit {
		target = left
	}

	if right < limit && list[right].profit > list[target].profit {
		target = right
	}

	if target != index {
		list[target], list[index] = list[index], list[target]
		adjustProfitHeap(list, target, limit)
	}
}
