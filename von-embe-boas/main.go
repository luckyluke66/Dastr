package main

import (
    "fmt"
    "math"
)

type VEBNode struct {
    u       int
    min     *int
    max     *int
    summary *VEBNode
    cluster []*VEBNode
}

type VEBTree struct {
    head *VEBNode
}

func high(x int, u int) int {
    return x / int(math.Sqrt(float64(u)))
}

func low(x int, u int) int {
    return x % int(math.Sqrt(float64(u)))
}

func Min(node VEBNode) *int {
    return node.min
}

func Max(node VEBNode) *int {
    return node.max
}

func Member(tree *VEBTree, val int) bool {
    var iter func(node *VEBNode, val int) bool

    iter = func(node *VEBNode, val int) bool {
        if node.min == nil || node.max == nil {
            return false
        }
        if *node.min == val || *node.max == val {
            return true
        }
        if node.u == 2 {
            return false
        }
        return iter(node.cluster[high(val, node.u)], low(val, node.u))
    }
    return iter(tree.head, val)
}

func Insert(tree *VEBTree, x int) {
    if tree.head.min == nil {
        tree.head.min = &x
        tree.head.max = &x
    } else {
        insert(tree.head, x)
    }
}

func insert(node *VEBNode, x int) {
    if node.min == nil {
        emptyInsert(node, x)
    } else if x < *node.min {
        x, *node.min = *node.min, x
    }
    if node.u > 2 {
        high, low := high(x, node.u), low(x, node.u)
        if node.cluster[high] == nil || node.cluster[high].min == nil {
            if node.summary != nil {
                insert(node.summary, high)
            }
            if node.cluster[high] != nil {
                emptyInsert(node.cluster[high], low)
            }
        } else {
            insert(node.cluster[high], low)
        }
    }
    if x > *node.max {
        node.max = &x
    }
}

func index(val int, offset int, u int) int {
    return val*int(math.Sqrt(float64(u))) + offset 
}

func successor(node *VEBNode, val int) *int {
    if node == nil || val >= *node.max {
        return nil
    }
    if val < *node.min {
        return node.min
    }
    if node.u > 2 {
        high, low := high(val, node.u), low(val, node.u)
        maxLow := node.cluster[high].max
        if maxLow != nil && low < *maxLow {
            offset := *successor(node.cluster[high], low)
            a := index(high, offset, node.u)
            return &a
        } else {
            succCluster := successor(node.summary, high)
            if succCluster == nil {
                return nil
            } else {
                offset := *node.cluster[*succCluster].min
                a := index(*succCluster, offset, node.u)
                return &a
            }
        }
    }
    return node.max
}


func emptyInsert(node *VEBNode, val int) {
    a := val
    node.min = &a
    node.max = &a
}

func printVEBTree(node *VEBNode, level int) {
    if node == nil {
        return
    }

    indent := ""
    for i := 0; i < level; i++ {
        indent += "  "
    }

    fmt.Printf("%sNode u: %d\n", indent, node.u)
    if node.min != nil {
        fmt.Printf("%sMin: %d\n", indent, *node.min)
    } else {
        fmt.Printf("%sMin: nil\n", indent)
    }
    if node.max != nil {
        fmt.Printf("%sMax: %d\n", indent, *node.max)
    } else {
        fmt.Printf("%sMax: nil\n", indent)
    }

    if node.u > 2 {
        if node.summary != nil {
            fmt.Printf("%sSummary:\n", indent)
            printVEBTree(node.summary, level+1)
        } else {
            fmt.Printf("%sSummary: nil\n", indent)
        }

        fmt.Printf("%sCluster:\n", indent)
        for i, clusterNode := range node.cluster {
            fmt.Printf("%sCluster %d:\n", indent, i)
            printVEBTree(clusterNode, level+1)
        }
    }
}

func printVEB(v *VEBTree) {
    printVEBTree(v.head, 0)
}

func NewWEBtree(u int) *VEBNode {
    if u < 2 {
        return nil
    }
    v := &VEBNode{u: u}
    if u > 2 {
        clusterSize := int(math.Ceil(math.Sqrt(float64(u))))
        v.summary = NewWEBtree(clusterSize)
        v.cluster = make([]*VEBNode, clusterSize)
        for i := 0; i < clusterSize; i++ {  
            v.cluster[i] = NewWEBtree(clusterSize)
        }
    }
    return v
}


func main() {
    //items := []int{2, 5, 6,7}
    tree := VEBTree{head: NewWEBtree(16)}

    Insert(&tree, 2)
    Insert(&tree, 5)
    Insert(&tree, 1)
    Insert(&tree, 6)
    Insert(&tree, 7)
    printVEB(&tree)

    fmt.Println(*successor(tree.head, 4))

    fmt.Println(Member(&tree, 2))
    fmt.Println(Member(&tree, 5))
}
