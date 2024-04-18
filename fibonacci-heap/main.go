package main

import (
    "fmt"
    "math"
)

type FibNode struct {
    key int
    left *FibNode
    right *FibNode
    child *FibNode
    parent *FibNode
    degree uint
    mark bool
}

type FibHeap struct {
    min *FibNode
    nodeCount uint 
}

func Insert(heap *FibHeap, key int) {
    node := &FibNode{key: key}
    node.left = node
    node.right = node
    heap.nodeCount++
    if heap.min == nil {
        heap.min = node
        return 
    }

    node.right = heap.min
    node.left = heap.min.left
    heap.min.left = node
    node.left.right = node

    if node.key < heap.min.key {
        heap.min = node
    }
}

func ExtractMin(heap *FibHeap) int {
    z := heap.min
    if z != nil {
        if z.child != nil {
            var count uint = 0
            child := z.child
            first := child
            start := child
            last := child
            
            for {
                child.parent = nil
                child = child.right
                if child == start {
                    break
                }
                last = child
                count++
            }

            last.right = z.right
            first.left = z.left
            z.left.right = first
            z.right.left = last
            heap.nodeCount += count
        } else {
            z.left.right = z.right
            z.right.left = z.left
        }
        heap.nodeCount--
        if z == z.right {
            heap.min = nil
        } else {
            heap.min = z.right
            Consolidate(heap)
        }
    }
    return z.key
}

func Consolidate(heap *FibHeap) {
    maxDegree := uint(math.Log2(float64(heap.nodeCount))) + 1
    arr := make([]*FibNode, maxDegree)
    w := heap.min
    last := w.left
    for {
        r := w.right
        x := w
        d := w.degree

        for arr[d] != nil {     //spadne protoze d je vetsi nez delka arraye
            y := arr[d]
            if x.key > y.key {
                x, y = y, x
            }
            FibHeapLink(heap, y, x)
            arr[d] = nil
            d++
        }
        arr[d] = x
        if w == last {
            break
        }
        w = r
    }
    heap.min = nil
    /*for _, v := range arr {
        if v != nil {
            Insert(heap, v)
        }
        
    }
    */
    for _ ,node := range arr {
        if node != nil {
            if heap.min == nil {
                heap.min = node
                node.left = node
                node.right = node
            } else {
                node.left = heap.min
                node.right = heap.min.right
                heap.min.right = node
                node.right.left = node
                if node.key < heap.min.key {
                    heap.min = node
                }
            }
        }
    }
}

func FibHeapLink(heap *FibHeap, y, x *FibNode) {
    y.left.right = y.right
    y.right.left = y.left
    
    if x.child == nil {
        x.child = y
        y.right = y
        y.left = y
    } else {
        y.left = x.child
        y.right = x.child.right
        y.left.right = y
        y.right.left = y
    }

    y.parent = x
    x.degree++
    y.mark = false
}

func DecreaseKey(heap *FibHeap, x *FibNode, k int) {
    if k > x.key {
        panic("new key is greater than current key")
    }
    x.key = k
    y := x.parent 
    if y != nil && x.key < y.key {
        Cut(heap, x, y)
        CascadingCut(heap, y)
    }
    if x.key < heap.min.key {
        heap.min = x
    }
}

func Cut(heap *FibHeap, x, y *FibNode) {
    if x == x.right {
        y.child = nil
    } else {
        x.left.right = x.right
        x.right.left = x.left
        if y.child == x {
            y.child = x.right
        }
    }
    y.degree--
    x.left = heap.min
    x.right = heap.min.right
    heap.min.right = x
    x.right.left = x
    x.parent = nil
    x.mark = false
}

func CascadingCut(heap *FibHeap, y *FibNode) {
    z := y.parent
    if z != nil {
        if !y.mark {
            y.mark = true
        } else {
            Cut(heap, y, z)
            CascadingCut(heap, z)
        }
    }
}
    
func printHeap(h *FibHeap, node *FibNode, prefix string) {
    if node == nil {
        return
    }

    fmt.Println(prefix, node.key)
    child := node.child
    if child != nil {
        printHeap(h, child, prefix+"  ")
        
        for c := child.right; c != child; c = c.right {
            printHeap(h, c, prefix+"  ")
        }
    }
}

func printFullHeap(h *FibHeap) {
    if h.min == nil {
        return
    }

    printHeap(h, h.min, "")
    
    for n := h.min.right; n != h.min; n = n.right {
        printHeap(h, n, "")
    }
}

func main() {
    h := &FibHeap{}
    
    Insert(h, 1)
    Insert(h, 2)
    Insert(h, 3)
    Insert(h, 4)
    printFullHeap(h)
    fmt.Println("extract")
    for i := 0; i < 4; i++ {
        fmt.Println(ExtractMin(h))
    }

    Insert(h, 100)
    DecreaseKey(h, h.min, 1)

    fmt.Println("extract")
    printFullHeap(h)
}