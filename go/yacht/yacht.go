package yacht

func timesScore(count map[int]int, num int) int {
	return num * count[num]
}

func fullHouseScore(count map[int]int) int {
	sum := 0
	prevTimes := 0
	for num, times := range count {
		if times == 2 || times == 3 {
			if prevTimes == times { // Seen again
				return 0
			}
			sum += times * num
			prevTimes = times
		}
	}
	return sum
}

func fourOfAKindScore(count map[int]int) int {
	for num, times := range count {
		if times >= 4 {
			return 4 * num
		}
	}
	return 0
}

func allOneTime(count map[int]int) bool {
	for _, times := range count {
		if times != 1 {
			return false
		}
	}
	return true
}

func littleStraightScore(count map[int]int) int {
	if count[1] != 1 || count[6] != 0 || !allOneTime(count) {
		return 0
	}
	return 30
}

func bigStraightScore(count map[int]int) int {
	if count[1] == 1 || !allOneTime(count) {
		return 0
	}
	return 30
}

func yachtScore(count map[int]int) int {
	for _, times := range count {
		if times == 5 {
			return 50
		}
	}
	return 0
}

func choiceScore(count map[int]int) int {
	sum := 0
	for num, times := range count {
		sum += times * num
	}
	return sum
}

func Score(dice []int, category string) int {
	count := map[int]int{}
	for _, num := range dice {
		count[num]++
	}
	switch category {
	case "ones":
		return timesScore(count, 1)
	case "twos":
		return timesScore(count, 2)
	case "threes":
		return timesScore(count, 3)
	case "fours":
		return timesScore(count, 4)
	case "fives":
		return timesScore(count, 5)
	case "sixes":
		return timesScore(count, 6)
	case "full house":
		return fullHouseScore(count)
	case "four of a kind":
		return fourOfAKindScore(count)
	case "little straight":
		return littleStraightScore(count)
	case "big straight":
		return bigStraightScore(count)
	case "choice":
		return choiceScore(count)
	case "yacht":
		return yachtScore(count)
	default:
		return 0
	}
}
