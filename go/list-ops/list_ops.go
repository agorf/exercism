package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, n := range s {
		acc = fn(acc, n)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filtered := IntList{}
	for _, n := range s {
		if fn(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	mapped := make(IntList, len(s))
	for i, n := range s {
		mapped[i] = fn(n)
	}
	return mapped
}

func (s IntList) Reverse() IntList {
	length := len(s)
	reversed := make(IntList, length)
	for i := range s {
		reversed[i] = s[length-i-1]
	}
	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	out := IntList{}
	out = append(out, s...)
	out = append(out, lst...)
	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	out := IntList{}
	out = append(out, s...)
	for _, list := range lists {
		out = append(out, list...)
	}
	return out
}
