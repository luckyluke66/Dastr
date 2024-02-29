package main

import (
	"fmt"
    "strings"
)

type BinomialNode struct {
    key     int
    degree  int
    child   *BinomialNode
    sibling *BinomialNode
    parent  *BinomialNode
}

type BinomialHeap struct {
    head *BinomialNode
}

func merge(heap1 *BinomialNode, heap2 *BinomialNode) *BinomialNode {
    if heap1 == nil {
        return heap2
    } else if heap2 == nil {
        return heap1
    } else {
        if heap1.key < heap2.key {
            heap1.sibling = merge(heap1.sibling, heap2)
            return heap1
        } else {
            heap2.sibling = merge(heap2.sibling, heap1)
            return heap2
        }
    }
}

func isCaseOne(x *BinomialNode, nextX *BinomialNode) bool {
    return (x.degree != nextX.degree) ||
        (nextX.sibling != nil && nextX.sibling.degree == x.degree)
}

func isCaseTwo(x *BinomialNode, nextX *BinomialNode) bool {
    return (x.key <= nextX.key)
}

func link(x *BinomialNode, y *BinomialNode) {
    x.parent = y
    x.sibling = y.child
    y.child = x
    y.degree++
}

func union(heap1 *BinomialHeap, heap2 *BinomialHeap) *BinomialHeap {
    h := merge(heap1.head, heap2.head)

    if h == nil {
        return nil
    }

    var prevX *BinomialNode = nil
    x := h
    nextX := x.sibling

    for nextX != nil {
        if isCaseOne(x, nextX) {
            prevX = x
            x = nextX
        } else if isCaseTwo(x, nextX) {
            x.sibling = nextX.sibling
            link(nextX, x)
        } else {
            if prevX == nil {
                h = nextX
            } else {
                prevX.sibling = nextX
            }
            link(x, nextX)
            x = nextX
        }
        nextX = x.sibling
    }
    return &BinomialHeap{h}
}

func insert(heap *BinomialHeap, key int) {
    newNode := BinomialNode{}
    newNode.key = key
    newHeap := &BinomialHeap{&newNode}
    *heap = *union(heap, newHeap)
}

func extractMin(heap *BinomialHeap) {
   //TODO: implement 
}

func decreaseKey(heap *BinomialHeap, node *BinomialNode, newKey int) {
    //TODO: implement
}


func printHeap(node *BinomialNode, level int) {
    if node == nil {
        return
    }

    fmt.Printf("%sNode: %d, Degree: %d\n", strings.Repeat("  ", level), node.key, node.degree)
    printHeap(node.child, level+1)
    printHeap(node.sibling, level)
}

func printBinomialHeap(heap *BinomialHeap) {
    if heap == nil || heap.head == nil {
        fmt.Println("Heap is empty")
        return
    }

    fmt.Println("Binomial Heap:")
    printHeap(heap.head, 0)
}

func main() {
    heap:= &BinomialHeap{}
    keys := []int{20, 50, 11, 7, 45, 3, 5, 6, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19}
    
    for _, key := range keys {
        insert(heap, key)
    }

    bl := heap.head == nil
    fmt.Println(bl)
    printBinomialHeap(heap)
}