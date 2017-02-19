package main

import (
	"fmt"
	"math/rand"
)

type Collection []int

func (self Collection) Partition(start int, end int) int {

	pindex := start
	pval := self[end-1]

	for i := start; i < end-1; i++ {
		if self[i] <= pval {
			self[i], self[pindex] = self[pindex], self[i]
			pindex++
		}
	}
	self[end-1], self[pindex] = self[pindex], self[end-1]
	return pindex
}

func (self Collection) Quicksort(start int, end int) {

	if start < end {
		pindex := self.Partition(start, end)
		self.Quicksort(start, pindex-1)
		self.Quicksort(pindex+1, end)
	}
}

func (self Collection) MergeSort() Collection {

	cLen := len(self)
	if cLen <= 2 {
		return self
	}

	ret := make(Collection, cLen)
	if cLen > 2 {

		half := len(self) / 2
		ll := self[0:half]
		lr := self[half:]

		ll = ll.MergeSort()
		lr = lr.MergeSort()

		k := 0
		for i, j := 0, 0; i < len(ll) || j < len(lr); {

			if i >= len(ll) {
				ret[k] = lr[j]
				j++
			} else if j >= len(lr) {
				ret[k] = ll[i]
				i++
			} else {
				c1 := ll[i]
				c2 := lr[j]

				if c1 < c2 {
					ret[k] = c1
					i++
				} else {
					ret[k] = c2
					j++
				}
			}
			k++
		}

	}

	return ret

}

func (self Collection) Print() {

	for _, elem := range self {
		fmt.Printf("%d ", elem)
	}
	fmt.Println("")

}

func main() {

	rand.Seed(42)
	list := make(Collection, 10)
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(100)
	}
	//list = list.MergeSort()
	list.Print()
	list.Quicksort(0, len(list))

	list.Print()

}
