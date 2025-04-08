package listops

type IntList []int

func (s IntList) Foldr(fn func(item, acc int) int, initial int) int {
	acc := initial
	for i := s.Length() - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Foldl(fn func(acc, item int) int, initial int) int {
	acc := initial
	for _, item := range s {
		acc = fn(acc, item)
	}
	return acc
}

func (s IntList) Filter(fn func(item int) bool) IntList {
	res := make(IntList, s.Length())
	var j int
	for _, item := range s {
		if fn(item) {
			res[j] = item
			j++
		}
	}
	return res[:j]
}

func (s IntList) Length() int {
	var count int
	for range s {
		count++
	}
	return count
}

func (s IntList) Map(fn func(item int) int) IntList {
	res := make(IntList, s.Length())
	for i := 0; i < s.Length(); i++ {
		res[i] = fn(s[i])
	}
	return res
}

func (s IntList) Reverse() IntList {
	n := s.Length()
	res := make(IntList, n)
	for i := 0; i < n; i++ {
		res[i] = s[n-1-i]
	}
	return res
}

func (s IntList) Append(list IntList) IntList {
	return s.Concat([]IntList{list})
}

func (s IntList) Concat(lists []IntList) IntList {
	totalLen := s.Length()
	for _, list := range lists {
		totalLen += list.Length()
	}

	res := make(IntList, totalLen)
	var i int
	for _, item := range s {
		res[i] = item
		i++
	}
	for _, list := range lists {
		for _, item := range list {
			res[i] = item
			i++
		}
	}
	return res
}
