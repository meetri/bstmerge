package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	//"log"
)

type BstNode struct {
	data int
	lf   *BstNode
	rt   *BstNode
}

func (self *BstNode) Depth() {
}

func (self *BstNode) Height() int {

	lHt, rHt := -1, -1
	if self.lf != nil {
		lHt = self.lf.Height()
	}

	if self.rt != nil {
		rHt = self.rt.Height()
	}

	max := rHt
	if lHt > rHt {
		max = lHt
	}

	return max + 1

}

func (self *BstNode) Search(d int) {
}

//insert node into tree
func (self *BstNode) Insert(d int) {

	if d < self.data {
		if self.lf == nil {
			self.lf = new(BstNode)
			self.lf.data = d
		} else {
			self.lf.Insert(d)
		}
	} else if d > self.data {
		if self.rt == nil {
			self.rt = new(BstNode)
			self.rt.data = d
		} else {
			self.rt.Insert(d)
		}

	}

}

//build sorted array into BST tree
func (self *BstNode) build(d []int) {

	mid := len(d) / 2
	dlen := len(d)
	self.data = d[mid]

	if dlen > 1 {
		self.lf = new(BstNode)
		self.rt = new(BstNode)
		self.lf.build(d[0:mid])
		self.rt.build(d[mid+1:])
	}
}

//Level Order Traversal
func (self *BstNode) lot(queue chan BstNode, processor func(data int)) {

	processor(self.data)
	if self.lf != nil {
		queue <- *self.lf
	}

	if self.rt != nil {
		queue <- *self.rt
	}

}

func (self *BstNode) Preorder(processor func(data int)) {

	processor(self.data)

	if self.lf != nil {
		self.lf.Preorder(processor)
	}

	if self.rt != nil {
		self.rt.Preorder(processor)
	}

}

func (self *BstNode) Inorder(processor func(data int)) {

	if self.lf != nil {
		self.lf.Inorder(processor)
	}

	processor(self.data)

	if self.rt != nil {
		self.rt.Inorder(processor)
	}

}

func (self *BstNode) Postorder(processor func(data int)) {

	if self.lf != nil {
		self.lf.Inorder(processor)
	}

	if self.rt != nil {
		self.rt.Inorder(processor)
	}

	processor(self.data)

}

func (self *BstNode) Levelorder(processor func(data int)) {

	queue := make(chan BstNode, 10)
	queue <- *self

	done := false
	for done == false {
		select {
		case node := <-queue:
			(&node).lot(queue, processor)
		default:
			done = true
		}
	}

}

func main() {

	nums := []int{8, 10, 12, 15, 17, 20, 25}
	rootPtr := new(BstNode)
	rootPtr.build(nums)

	rootPtr.Insert(22)
	rootPtr.Insert(3)
	rootPtr.Insert(9)
	rootPtr.Insert(28)
	rootPtr.Insert(32)

	printProcessor := func(data int) {
		fmt.Printf("%d,", data)
	}

	fmt.Printf("Preorder: ")
	rootPtr.Preorder(printProcessor)

	fmt.Printf("\nInorder: ")
	rootPtr.Inorder(printProcessor)

	fmt.Printf("\nPostorder: ")
	rootPtr.Postorder(printProcessor)

	fmt.Printf("\nLevel-order traversal: ")
	rootPtr.Levelorder(printProcessor)

}
