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
    }
    if x < *node.min {
        x, *node.min = *node.min, x
    }
    if node.u > 2 {
        high, low := high(x, node.u), low(x, node.u)
        if node.cluster[high].min == nil {
            fmt.Println(node.summary, high)
            insert(node.summary, high)
            emptyInsert(node.cluster[high], low)
        } else {
            insert(node.cluster[high], low)
        }
    }
    if x > *node.max {
        node.max = &x
    }
}

func emptyInsert(node *VEBNode, val int) {
    a := val
    node.min = &a
    node.max = &a
}

func MkVEBNode(u int) *VEBNode {
    node := &VEBNode{}
    node.u = u
    if u > 2 {
        sqrt := int(math.Sqrt(float64(u)))
        node.cluster = make([]*VEBNode, sqrt)
        for i := 0; i < sqrt; i++ {
            node.cluster[i] = MkVEBNode(sqrt)
        }
    }
    return node
}

func MkVEBTree(u int) VEBTree {
    tree := VEBTree{}
    tree.head = MkVEBNode(u)
    tree.head.summary = MkSummary(tree.head.u / 2)
    return tree
}

func MkSummary(u int) *VEBNode {
    summary := &VEBNode{}
    div := u / 2
    summary.u = div
    if u > 2 {
        summary.cluster = make([]*VEBNode, div)
        for i := 0; i < div; i++ {
            summary.cluster[i] = MkSummary(div)
        }

        summary.summary = MkSummary(div)
    }
    return summary
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

func main() {
    //items := []int{2, 5, 6,7}
    tree := MkVEBTree(16)

    Insert(&tree, 2)
    Insert(&tree, 5)
    //Insert(tree.head, 6)
    //Insert(&tree, 7)
    printVEB(&tree)

    fmt.Println(Member(&tree, 2))
    fmt.Println(Member(&tree, 5))
}
