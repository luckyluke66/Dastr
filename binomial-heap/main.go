package main

import (
	"fmt"
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
    }
    if heap2 == nil {
        return heap1
    }

    var head *BinomialNode

    if heap1.degree < heap2.degree {
        head = heap1
        heap1 = heap1.sibling
    } else {
        head = heap2
        heap2 = heap2.sibling
    }

    tail := head

    for heap1 != nil && heap2 != nil {
        if heap1.degree < heap2.degree {
            tail.sibling = heap1
            heap1 = heap1.sibling
        } else {
            tail.sibling = heap2
            heap2 = heap2.sibling
        }
        tail = tail.sibling
    }

    if heap1 != nil {
        tail.sibling = heap1
    } else {
        tail.sibling = heap2
    }

    return head
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

func removeMinNode(heap *BinomialHeap) *BinomialNode {
    if heap.head == nil {
        return nil
    }

    prevX := heap.head
    x := heap.head
    nextX := x.sibling
    min := x
    prevMin := prevX

    for nextX != nil {
        if nextX.key < min.key {
            min = nextX
            prevMin = prevX
        }
        prevX = nextX
        nextX = nextX.sibling
    }

    if prevMin == prevX {
        heap.head = x.sibling
    } else {
        prevMin.sibling = min.sibling
    }

    return min
}

func reverseList(node *BinomialNode) *BinomialNode {
    if node == nil || node.sibling == nil {
       return node
    }
    rest := reverseList(node.sibling)
    node.sibling.sibling = node
    node.sibling = nil
    return rest
}

func extractMin(heap *BinomialHeap) int {
    x := removeMinNode(heap)
    y := reverseList(heap.head.child)

    heap = union(heap, &BinomialHeap{y})
    return x.key
}

func decreaseKey(heap *BinomialHeap, node *BinomialNode, newKey int) {
    if newKey > node.key {
        return
    }

    node.key = newKey
    y := node
    z := y.parent

    for z != nil && y.key < z.key {
        y.key, z.key = z.key, y.key
        y = z
        z = y.parent
    }
}

func printHeapHelper(node *BinomialNode, level int, nodeType string) {
    if node == nil {
        return
    }

    fmt.Printf("Level: %d, Key: %d, Node Type: %s\n", level, node.key, nodeType)
    printHeapHelper(node.child, level+1, "Child")
    printHeapHelper(node.sibling, level, "Sibling")
}

func printHeap(heap *BinomialHeap) {
    printHeapHelper(heap.head, 0, "Head")
}

func main() {
    heap:= &BinomialHeap{}
    keys := []int{20, 50, 11, 7, 45, 3, 5, 6, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19}
    
    for _, key := range keys {
        insert(heap, key)
    }

    printHeap(heap)
}