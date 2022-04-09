package main

type Items []*Account

func (h Items) Len() int {
	return len(h)
}

func (h Items) Less(i, j int) bool {
	return h[i].nextAccessTime.Before(h[j].nextAccessTime)
}

func (h Items) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Items) Push(item interface{}) {
	*h = append(*h, item.(*Account))
}

func (h *Items) Pop() interface{} {
	before := *h
	n := len(before) - 1
	item := before[n]
	*h = before[0:n]
	return item
}
