package list

type CostProfit struct {
	cost, profit int
}

func MaxMoney(cost, profit []int, money, choice int) int {
	costList := make([]*CostProfit, 0)
	for i, value := range cost {
		costList = append(costList, &CostProfit{
			cost:   value,
			profit: profit[i],
		})
	}
	initCostHeap(costList)

	ans := 0
	for choice > 0 && money > 0 {
		getCostHeap(costList)
	}
	return ans

}

func initCostHeap(list []*CostProfit) {
	begin := len(list)/2 - 1
	for i := begin; i < len(list); i++ {
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

func getCostHeap(list []*CostProfit) (int, []*CostProfit) {
	value := list[0]
	list[0] = list[len(list)-1]
	list = list[:len(list)-1]
	adjustCostHeap(list, 0, len(list)-1)
	return value.cost, list
}

func initProfitHeap(list []*CostProfit) {
	begin := len(list)/2 - 1
	for i := begin; i < len(list); i++ {
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

func getProfitHeap(list []*CostProfit) (int, []*CostProfit) {
	value := list[0]
	list[0] = list[len(list)-1]
	list = list[:len(list)-1]
	adjustProfitHeap(list, 0, len(list)-1)
	return value.profit, list
}
